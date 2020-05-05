import React, { Component } from 'react';
import Content from "../../components/Content";
import {Box, Button, Col, Inputs, Row} from "adminlte-2-react";
import WorkTable from "../../components/WorkTable";
import moment from "moment";

const { Select2, DateRange } = Inputs;
const uuid = require('uuid/v4');

class OrderCreate extends Component {
    constructor(props) {
        super(props);
        this.state = {
            startDate: null,
            endDate: null,
            focusedInput: null,
            order: {
                quoteid: parseInt((new URLSearchParams(this.props.location.search)).get('quoteid')),
                employees: [],
                detail:[{idx: uuid(), workhours: 0, cloths: 0, description: '', subdetail: []}],
                aprovedate: '',
                deliverydate: '',
                collaborators: '',
                totalhours: 0,
                totalcloths: 0,
                startdate: moment().add(1, 'days').format('DD/MM/YYYY'),
                enddate: moment().add(10, 'days').format('DD/MM/YYYY')
            },
            customers: [],
            insumos: [],
            services: [],
            products: [],
            employees: [],
            quote: [],
            supplies: []
        };
        this.handleSubmit = this.handleSubmit.bind(this);
    }
    ws = new WebSocket('ws://localhost:8080/ws');
    handleSubmit(event) {
        if (window.confirm("Esta seguro?")) {
            fetch('http://localhost:8080/api/v1/orders', {
                method: 'POST',
                body: JSON.stringify(this.state.order)
            })
            .then(d=>d.json())
            .then(()=> {
                this.ws.send("create order");
                this.ws.close();
                this.props.history.push("/order")
            })
            .catch(console.log)
        } else {
            return false
        }
    }
    componentDidMount() {
        const search = this.props.location.search;
        const params = new URLSearchParams(search);
        const quoteid = parseInt(params.get('quoteid'));
        if (quoteid) {
            fetch('http://localhost:8080/api/v1/supplies/quote/' + quoteid)
                .then(res => res.json())
                .then(data => {
                    this.setState({supplies: data})
                })
                .catch(console.log);
        }
        fetch('http://localhost:8080/api/v1/employees')
            .then(res => res.json())
            .then(data => {this.setState({employees: data})})
            .catch(console.log);
        fetch('http://localhost:8080/api/v1/quotations')
            .then(res => res.json())
            .then(data => {this.setState({quote: data})})
            .catch(console.log);
        fetch('http://localhost:8080/api/v1/products')
            .then(res => res.json())
            .then(data => {
                this.setState({
                    products: data,
                    services: data.filter(q=>q.category==='SERVICIO'),
                    insumos: data.filter(q=>q.category==='INSUMOS')
                })
            })
            .catch(console.log);
    }
    render() {
        const handleChange = e => {
            let value = e.target.value;
            let subid = e.target.name.split('-').length > 2 ? e.target.name.split('-')[2] : e.target.id;
            let id = e.target.name.split('-').length > 1 ? e.target.name.split('-')[1] : e.target.id;
            let name = e.target.name.split('-').length > 0 ? e.target.name.split('-')[0] : e.target.name;
            let order = {...this.state.order};
            if (id==='startdate') {
                order.startdate = value;
            }
            if (id==='enddate') {
                order.enddate = value;
            }
            if (name==='quoteid') {
                order.quoteid = parseInt(value);
            }
            if (name==='worktype') {
                order.worktype = value;
            }
            /*if (['employee1', 'employee2'].includes(name)) {
                let options = e.target.options;
                for (let i = 0, l = options.length; i < l; i++) {
                    let v = parseInt(options[i].value);
                    if (options[i].selected && !order['employees'].includes(v)) {
                        order['employees'].push(v);
                    }
                }
            }*/
            if (name === 'description') {
                order.detail[id]["description"]=value;
            }
            if (name === 'employee') {
                order.detail[id].subdetail[subid].employee=parseInt(value);
            }
            if (name === 'subworkhours') {
                order.detail[id].subdetail[subid].workhours=parseInt(value);
                order.detail[id].workhours = order.detail[id].subdetail?order.detail[id].subdetail.map(d=>d.workhours).reduce((a,b)=>a+b):0
                order.totalhours = order.detail.map(d=>d.workhours).reduce((a,b)=>a+b,0)
            }
            if (name === 'subcloths') {
                order.detail[id].subdetail[subid].cloths=parseInt(value);
                order.detail[id].cloths = order.detail[id].subdetail?order.detail[id].subdetail.map(d=>d.cloths).reduce((a,b)=>a+b):0
                order.totalcloths = order.detail.map(d=>d.cloths).reduce((a,b)=>a+b,0)
            }
            this.setState({order: order});
        };
        const handleAddItem = () => {
            let detail = new Object();
            detail.idx = uuid();
            detail.workhours = 0;
            detail.cloths = 0;
            detail.description = '';
            detail.subdetail = [];
            let order = {...this.state.order};
            order.detail = [...this.state.order.detail, detail];
            this.setState({order: order});
        };
        const handleDelItem = idx => {
            let order = {...this.state.order};
            order.detail = order.detail.filter(detail=>detail.idx !== idx);
            this.setState({
                order: order
            })
        };
        const handleAddSubitem = idx => {
            let order = {...this.state.order};
            let id = order.detail.findIndex(d=>d.idx===idx);
            let subdetail = new Object();
            subdetail.idx = uuid();
            subdetail.workhours = 0;
            subdetail.cloths = 0;
            order.detail[id].subdetail = [...order.detail[id].subdetail, subdetail];
            this.setState({
                order: order
            },()=>console.log(this.state.order));
        }
        const handleDelSubitem = (idx, subidx) => {
            let order = {...this.state.order};
            let id = order.detail.findIndex(d=>d.idx===idx);
            order.detail[id].subdetail = order.detail[id].subdetail.filter(subdetail=>subdetail.idx !== subidx);
            this.setState({
                order: order
            })
        };
        const validRenderComplete = () => {
            return true
                && this.state.employees.length > 0
                && this.state.insumos.length> 0
                && this.state.services.length> 0
                && true;
        };
        const ButtonCreateOrder = () => {
            return <Button
                type={'danger'}
                text={'Guardar'}
                onClick={this.handleSubmit}
                pullRight
            />;
        };
        return (<Content title={this.props.title} subTitle={'Crear'} breadCrumb={[
            {title: "Orden de trabajo", link: "/order"},
            {title: "Nuevo"},
        ]} loaded={validRenderComplete()}>
                <Box>
                    <Row>
                        <Col md={4}>
                            <Select2
                                name={'quoteid'}
                                label={'Cotizacion'}
                                options={this.state.quote.map(p=> p.id)}
                                value={this.state.order.quoteid}
                                labelPosition={'above'}
                                onChange={handleChange}
                            />
                        </Col>
                        <Col md={4}>
                            <Select2
                                name={'worktype'}
                                label={'Tipo de trabajo'}
                                options={['PDI', 'SINIESTRO', 'UNIDAD COMPLETA', 'PIEZAS']}
                                value={this.state.order.worktype}
                                labelPosition={'above'}
                                onChange={handleChange}
                            />
                        </Col>
                        <Col md={4}>
                            <DateRange
                                label={'Fecha Inicio Entrega'}
                                labelPosition={'above'}
                                startDateId={'startdate'}
                                endDateId={'enddate'}
                                startDate={this.state.order.startdate}
                                endDate={this.state.order.enddate}
                                onStartChange={handleChange}
                                onEndChange={handleChange}
                                format={'DD/MM/YYYY'}
                            />
                        </Col>
                    </Row>

                </Box>
                <Box collapsable title={'Descripcion del trabajo'} footer={<ButtonCreateOrder/>}>
                    <Row>
                        <Col xs={12}>
                            <WorkTable
                                details={this.state.order.detail}
                                handleAddItem={handleAddItem}
                                handleDelItem={handleDelItem}
                                handleAddSubitem={handleAddSubitem}
                                handleDelSubitem={handleDelSubitem}
                                handleChange={handleChange}
                                employee={this.state.employees.filter(e=>e.care==="COLABORADOR").map(e=>{return {value: e.id, label: e.name + ' ' + e.lastname}})}
                            />
                        </Col>
                    </Row>
                </Box>
        </Content>)
    }
}

export default OrderCreate;