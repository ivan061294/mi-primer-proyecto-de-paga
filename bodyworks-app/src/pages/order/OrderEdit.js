import React, { Component } from 'react';
import Content from "../../components/Content";
import {Box, Col, Inputs, Row, Button } from "adminlte-2-react";
import moment from "moment";
import WorkTable from "../../components/WorkTable";

const { Select2, DateRange } = Inputs;
const uuid = require('uuid/v4');

class OrderEdit extends Component {
    constructor(props) {
        super(props);
        this.state = {
            order: {
                detail:[{idx: uuid(), workhours: 0, cloths: 0, description: '', subdetail: []}],
            },
            quote: [],
            products: [],
            employees: [],
            employees1: [],
            employees2: []
        };
        this.handleSubmit = this.handleSubmit.bind(this);
    }
    handleSubmit(event) {
        if (window.confirm("Estas seguro?")) {
            fetch('http://localhost:8080/api/v1/orders/'+ this.props.match.params.id, {
                method: 'PUT',
                body: JSON.stringify(this.state.order)
            })
                .then(this.props.history.push("/order"))
                .catch(console.log)
        } else {
            return false
        }
    }
    componentDidMount() {
        fetch('http://localhost:8080/api/v1/orders/' + this.props.match.params.id)
            .then(res => res.json())
            .then(data => {
                data.startdate = data.startdate?moment(data.startdate).format('DD/MM/YYYY'):null;
                data.enddate = data.enddate?moment(data.enddate).format('DD/MM/YYYY'):null;
                this.setState({order: data})
            })
            .catch(console.log);
        fetch('http://localhost:8080/api/v1/products')
            .then(res => res.json())
            .then(data => {
                this.setState({products: data})
            })
            .catch(console.log);
        fetch('http://localhost:8080/api/v1/employees')
            .then(res => res.json())
            .then(data => {this.setState({employees: data})})
            .catch(console.log);
        fetch('http://localhost:8080/api/v1/quotations')
            .then(res => res.json())
            .then(data => {
                this.setState({quote: data})
            })
            .catch(console.log);
    }
    render() {
        let optionStatus = [
            {value:'E', label:'En proceso'},
            {value:'T', label:'Terminado'},
            {value:'C', label:'Cancelado'}
        ];
        const handleChange = e => {
            let value = e.target.value;
            let id = e.target.name.split('-').length > 1 ? e.target.name.split('-')[1] : e.target.id;
            let name = e.target.name.split('-').length > 0 ? e.target.name.split('-')[0] : e.target.name;
            let order = {...this.state.order};
            if (name==='status') {
                order.status = value;
            }
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
            if (['employee1', 'employee2'].includes(name)) {
                let options = e.target.options;
                order[name] = [];
                for (let i = 0, l = options.length; i < l; i++) {
                    if (options[i].selected) {
                        order[name].push(parseInt(options[i].value));
                    }
                }
            }
            if (name === 'description') {
                order.detail[id]["description"]=value;
            }
            if (name === 'workhours') {
                order.detail[id]["workhours"]=parseInt(value);
                order.totalhours = order.detail?order.detail.map(d=>d.workhours).reduce((a,b)=>a+b):0
            }
            if (name === 'cloths') {
                order.detail[id]["cloths"]=parseInt(value);
                order.totalcloths = order.detail?order.detail.map(d=>d.cloths).reduce((a,b)=>a+b):0
            }
            this.setState({order: order});
        };
        const handleAddItem = () => {
            let detail = new Object();
            detail.idx = uuid();
            detail.id = null;
            detail.quantity = 1;
            detail.product = 0;
            detail.description = '';
            detail.unitprice = 0.00;
            detail.price = 0.00;
            let order = {...this.state.order};
            order.detail = [...this.state.order.detail, detail];
            this.setState({
                order: order
            });
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
            console.log(order.detail[id])
            order.detail[id].subdetail = [...order.detail[id].subdetail, subdetail];
            this.setState({order: order});
        }
        const handleDelSubitem = (idx, subidx) => {
            let order = {...this.state.order};
            let id = order.detail.findIndex(d=>d.idx===idx);
            order.detail[id].subdetail = order.detail[id].subdetail.filter(subdetail=>subdetail.idx !== subidx);
            this.setState({order: order})
        };
        const validRenderComplete = () => {
            return true
                && this.state.order.id >= 0
                && this.state.employees.length > 0
                && this.state.products.length > 0
                && this.state.quote.length > 0
                && true;
        };

        const isSupervisor = (element) => {
            return this.state.employees.find(e=>e.id===element && e.care==="SUPERVISOR")
        };
        const isColaborate = (element) => {
            return this.state.employees.find(e=>e.id===element && e.care==="COLABORADOR")
        };
        return (<Content title={this.props.title} subTitle={'Modificación'} breadCrumb={[
            {title: "Orden de trabajo", link: "/order"},
            {title: this.props.match.params.id},
        ]} loaded={validRenderComplete()}>
                <Box collapsable title={'Principal'}>
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
                        <Col md={3}>
                            <Select2
                                name={'worktype'}
                                label={'Tipo de trabajo'}
                                options={['PDI', 'SINIESTRO', 'UNIDAD COMPLETA', 'PIEZAS']}
                                value={this.state.order.worktype}
                                labelPosition={'above'}
                                onChange={handleChange}
                            />
                        </Col>
                        <Col md={3}>
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
                        <Col md={2}>
                            <Select2
                                name={'status'}
                                labelPosition={'above'}
                                label={'Estado'}
                                options={optionStatus}
                                value={this.state.order.status}
                                onChange={handleChange}
                            />
                        </Col>
                    </Row>
                    <Row>
                        <Col md={12}>
                            <Select2
                                name={'employee1'}
                                labelPosition={'above'}
                                label={'Encargado del trabajo'}
                                options={this.state.employees.map(e=>{return {value: e.id, label: e.name+' '+e.lastname}})}
                                multiple
                                value={this.state.employees1.length>0?this.state.employees1:this.state.order.employees?this.state.order.employees.filter(isSupervisor):null}
                                onChange={handleChange}
                            />
                            <Select2
                                name={'employee2'}
                                labelPosition={'above'}
                                label={'Personal de apoyo'}
                                options={this.state.employees.map(e=>{return {value: e.id, label: e.name+' '+e.lastname}})}
                                multiple
                                value={this.state.order.employees?this.state.order.employees.filter(isColaborate):null}
                                onChange={handleChange}
                            />
                        </Col>
                    </Row>
                </Box>
                <Box collapsable title={'Descripcion del trabajo'}>
                    <Row>
                        <Col xs={12}>
                            <WorkTable
                                details={this.state.order.detail}
                                handleAddItem={handleAddItem}
                                handleDelItem={handleDelItem}
                                handleChange={handleChange}
                                handleAddSubitem={handleAddSubitem}
                                handleDelSubitem={handleDelSubitem}
                                employee={this.state.employees.filter(e=>e.care==="COLABORADOR").map(e=>{return {value: e.id, label: e.name + ' ' + e.lastname}})}
                            />
                        </Col>
                    </Row>
                    <Row>
                        <Col xs={12}>
                            <Button
                                type={'danger'}
                                text={'Actualizar'}
                                onClick={this.handleSubmit}
                                pullRight
                            />
                        </Col>
                    </Row>
                </Box>
        </Content>)
    }
}

export default OrderEdit;