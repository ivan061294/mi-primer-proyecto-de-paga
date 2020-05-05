import React, { Component } from 'react';
import Content from "../../components/Content";
import {Box, Col, Inputs, Row, Button } from "adminlte-2-react";
import SaleTable2 from "../../components/SaleTable2";

const { Text, Select2 } = Inputs;
const uuid = require('uuid/v4');

class QuoteEdit extends Component {
    constructor(props) {
        super(props);
        this.state = {
            quotation: {},
            customers: [],
            products: []
        };
        this.handleSubmit = this.handleSubmit.bind(this);
    }
    handleSubmit(event) {
        if (window.confirm("Estas seguro?")) {
            fetch('http://localhost:8080/api/v1/quotations/'+ this.props.match.params.id, {
                method: 'PUT',
                body: JSON.stringify(this.state.quotation)
            })
                .then(()=> this.props.history.push("/quote"))
                .catch(console.log)
        } else {
            return false
        }
    }
    componentDidMount() {
        fetch('http://localhost:8080/api/v1/quotations/' + this.props.match.params.id)
            .then(res => res.json())
            .then(data => {
                this.setState({quotation: data})
            })
            .catch(console.log);
        fetch('http://localhost:8080/api/v1/products')
            .then(res => res.json())
            .then(data => {
                this.setState({products: data})
            })
            .catch(console.log);
        fetch('http://localhost:8080/api/v1/customers')
            .then(res => res.json())
            .then(data => {
                this.setState({customers: data})
            })
            .catch(console.log);
    }
    render() {
        let optionStatus = [
            {value:'A', label:'Aceptado'},
            {value:'P', label:'Pendiente'},
            {value:'R', label:'Rechazado'}
        ];
        const handleChange = e => {
            let id = e.target.name.split('-')[1];
            let name = e.target.name.split('-').length>0 ? e.target.name.split('-')[0] : e.target.name;
            let value = e.target.value;
            let quotation = {...this.state.quotation};
            if (name === 'brand') {
                quotation.brand = value;
            }
            if (name === 'model') {
                quotation.model = value;
            }
            if (name === 'plate') {
                quotation.plate = value;
            }
            if (name === 'serie') {
                quotation.serie = value;
            }
            if (name === 'currency') {
                quotation.currency = value;
            }
            if (name === 'customer') {
                quotation.customer = parseInt(value);
            }
            if (name === 'status') {
                quotation.status = value;
            }
            if (name === 'contact') {
                quotation.contact = value;
            }
            if (name === 'product') {
                quotation.detail[id]["product"]=parseInt(value);
                quotation.detail[id]["unitprice"]=this.state.products.find(p=>p.id===parseInt(value)).unitprice;
            }
            if (name === 'description') {
                quotation.detail[id]["description"]=value;
            }
            if (name === 'quantity') {
                quotation.detail[id]["quantity"]=parseInt(value);
            }
            if (name === 'unitprice') {
                quotation.detail[id]["unitprice"]=parseFloat(value);
            }
            quotation.detail = quotation.detail.map(detail => {
                detail.price = detail.quantity * detail.unitprice;
                return detail
            });
            if (quotation.detail.length > 1) {
                quotation.total = quotation.detail.map(q=>q.price).reduce((a, b) => a + b) * 1.18;
            } else
                quotation.total = quotation.detail[0].price * 1.18;
            this.setState({
                quotation: quotation
            });
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
            let quotation = {...this.state.quotation};
            quotation.detail = [...this.state.quotation.detail, detail];
            this.setState({
                quotation: quotation
            });
        };
        const handleDelItem = idx => {
            let quotation = {...this.state.quotation};
            quotation.detail = quotation.detail.filter(detail=>detail.idx !== idx);
            this.setState({
                quotation: quotation
            })
        };
        const validRenderComplete = () => {
            return true
                && this.state.quotation.id>=0
                && this.state.products.length>0
                && this.state.customers.length>0
                && true;
        };
        const ButtonUpdate = () => {
            return <Button
                type={'danger'}
                text={'Actualizar'}
                onClick={this.handleSubmit}
                pullRight
            />;
        };
        return (<Content title={this.props.title} subTitle={'Modificación'} breadCrumb={[
            {title: "Cotizacion", link: "/quote"},
            {title: this.props.match.params.id},
        ]} loaded={validRenderComplete()}>
            <Box title={'Datos Generales'} footer={<ButtonUpdate/>}>
                <Row>
                    <Col md={4}>
                        <Select2
                            name={'customer'}
                            label={'Cliente'}
                            options={this.state.customers.map(p=>{return {value:p.id, label:p.fullname}})}
                            value={this.state.quotation.customer}
                            labelPosition={'above'}
                            onChange={handleChange}
                        />
                        <Text
                            name={'contact'}
                            labelPosition={'above'}
                            label={'Contacto'}
                            value={this.state.quotation.contact}
                            onChange={handleChange}
                        />
                    </Col>
                    <Col md={2}>
                        <Text
                            name={'color'}
                            labelPosition={'above'}
                            label={'Color'}
                            value={this.state.quotation.color}
                            onChange={handleChange}
                        />
                    </Col>
                    <Col md={2}>
                        <Text
                            name={'brand'}
                            labelPosition={'above'}
                            label={'Marca'}
                            value={this.state.quotation.brand}
                            onChange={handleChange}
                        />
                        <Text
                            name={'model'}
                            labelPosition={'above'}
                            label={'Modelo'}
                            value={this.state.quotation.model}
                            onChange={handleChange}
                        />
                    </Col>
                    <Col md={2}>
                        <Text
                            name={'plate'}
                            labelPosition={'above'}
                            label={'Placa'}
                            value={this.state.quotation.plate}
                            onChange={handleChange}
                        />
                        <Text
                            name={'serie'}
                            labelPosition={'above'}
                            label={'Serie'}
                            value={this.state.quotation.serie}
                            onChange={handleChange}
                        />
                    </Col>
                    
                    <Col md={2}>
                        <Select2
                            name={'status'}
                            labelPosition={'above'}
                            label={'Estado'}
                            options={optionStatus}
                            value={this.state.quotation.status}
                            optionState={this.state.status}
                            onChange={handleChange}
                        />
                    </Col>
                </Row>
                <hr/>
                <Row>
                    <Col xs={12}>
                        <h4>Carroceria y Pintura</h4>
                    </Col>
                </Row>
                <Row>
                    <Col xs={12}>
                        <SaleTable2
                            details={this.state.quotation.detail}
                            handleAddItem={handleAddItem}
                            handleDelItem={handleDelItem}
                            handleChange={handleChange}
                            listProducts={this.state.products.map(p=>{return {value:p.id, label:p.description}})}
                            currency={this.state.quotation.currency}
                        />
                    </Col>
                </Row>
            </Box>
        </Content>)
    }
}

export default QuoteEdit;