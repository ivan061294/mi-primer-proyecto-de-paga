import React, {Component} from 'react';

import {Box, Col, Inputs, Row, Button} from "adminlte-2-react";
import Content from "../../components/Content";
import SaleTable2 from "../../components/SaleTable2";

const { Text, Select2 } = Inputs;
const uuid = require('uuid/v4');

const zero = parseFloat(0).toFixed(2);

class InvoiceCreate extends Component {
    constructor(props) {
        super(props);
        this.state = {
            customers:[],
            invoice: {details:[{id: uuid(), quantity: 1, product: null, unitprice: zero, price: zero}]},
            saleprice: zero,
            igv: zero,
            total: zero,
            products:[],
            customer: null,
            contact: '',
            currency: 'PEN'
        };
        this.state.pivot = [
            {title: 'Valor de venta', value: this.state.saleprice},
            {title: 'IGV', value: this.state.igv},
            {title: 'Total', value: this.state.total}
        ];
        this.columns = [
            {title: 'Accion', width: 5, data: 'id', render: (id) => (<Button
                    icon={'fa-trash'}
                    type={'danger'}
                    onClick={this.handleDelItem.bind(this, id)}
                />)},
            {title: 'Cantidad', width: 10, data: 'quantity', render: (e, idx) => (<input
                    name={'quantity-' + idx}
                    value={this.state.details[idx].quantity}
                    className={'form-control'}
                    type={'number'}
                    onChange={this.handleChange}
                    min={1}
                />)},
            {title: 'Descripcion', width: 400, data: 'product', render: (e, index) => (<Select2
                    name={'product-' + index}
                    labelPosition={'none'}
                    options={this.state.products.map(p=>{return {value:p.id, label:p.description}})}
                    onChange={this.handleChange}
                />)},
            {title: 'Precio Unitario', align: 'right', width: 20, data: 'unitprice'},
            {title: 'Precio total', width: 20, align: 'right', data: 'price'},
        ];
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleAddItem = this.handleAddItem.bind(this);
        this.handleDelItem = this.handleDelItem.bind(this);
    }
    handleAddItem(event) {
        let detail = new Object();
            detail.idx = uuid();
            detail.id = null;
            detail.quantity = 1;
            detail.product = 0;
            detail.description = '';
            detail.unitprice = 0.00;
            detail.price = 0.00;
        let invoice = {...this.state.invoice};
        invoice.details = [...this.state.invoice.details, detail];
        this.setState({
            invoice: invoice
        });
    }
    handleDelItem(id) {
        this.setState({
            details: this.state.details.filter(detail=>detail.id !== id)
        })
    }
    handleSubmit(event) {
        if (window.confirm("Estas seguro?")) {
            fetch('http://localhost:8080/api/v1/invoices', {
                method: 'POST',
                body: JSON.stringify({
                    employee: 1,
                    customer: Number(this.state.customer),
                    contact: this.state.contact,
                    total: this.state.total,
                    detail: this.state.details.map(p => {
                        return {
                            product: Number(p.product),
                            amount: Number(p.quantity),
                            price: parseFloat(p.unitprice)
                        }
                    }),
                })
            })
                .then(res => {
                    console.log(res.json())
                    this.props.history.push("/quote")
                })
                .catch(console.log)
        } else {
            return false
        }
    }
    handleChange(event) {
        let id = event.target.name.split('-')[1];
        let name = event.target.name.split('-')[0];
        if (["quantity" ,"product", "unitprice"].includes(name)) {
            let details = [...this.state.details]
            details[id][name] = event.target.value
            if ("product" === name) {
                details[id]["unitprice"] = parseFloat(
                    this.state.products.find(d=>d.id===parseInt(event.target.value)).unitprice
                ).toFixed(2)
            }
            let saleprice = this.state.details.map(d => d.unitprice * d.quantity).reduce((a, b) => a + b)
            let igv = saleprice * 0.18;
            let total = saleprice + igv;
            let pivot = [];
            pivot.push({title: 'Valor de venta', value: parseFloat(saleprice).toFixed(2)});
            pivot.push({title: 'IGV', value: parseFloat(igv).toFixed(2)});
            pivot.push({title: 'Total', value: parseFloat(total).toFixed(2)});
            this.setState({
                details: details.map( d => {d.price = parseFloat(d.quantity * d.unitprice).toFixed(2); return d}),
                pivot: pivot,
                total: parseFloat(total).toFixed(2)
            })
        } else {
            const target = event.target;
            const value = target.type === 'checkbox' ? target.checked : target.value;
            const name = target.name;
            this.setState({
                [name]: value
            });
        }
    }
    componentDidMount() {
        const search = this.props.location.search;
        const params = new URLSearchParams(search);
        const quoteid = parseInt(params.get('quoteid'));
        if (quoteid) {
        fetch('http://localhost:8080/api/v1/quotations/' + quoteid)
            .then(res => res.json())
            .then(data => {
                this.setState({
                    invoice: data
                })
            })
            .catch(console.log);
        }

        fetch('http://localhost:8080/api/v1/customers')
            .then(res => res.json())
            .then(data => {
                this.setState({customers: data})
            })
            .catch(console.log);
        fetch('http://localhost:8080/api/v1/products')
            .then(res => res.json())
            .then(data => {
                this.setState({products: data})
            })
            .catch(console.log);
    }
    render() {
        const validRenderComplete = () => {
            return true
            && this.state.customers.length > 0
            && this.state.products.length > 0
            && true;
        };
        console.log(this.state.invoice.customer)
        return (<Content title={this.props.title} subTitle={'Crear'} breadCrumb={[
            {title: "Orden de trabajo", link: "/order"},
            {title: "Nuevo"},
        ]} loaded={validRenderComplete()}>
            <Row>
                <Col xs={12}>
                    <Box>
                        <Col md={4}>
                            <Select2
                                label="Cliente"
                                options={this.state.customers.map(p=>{return {value:p.id, label:p.fullname}})}
                                name={'customer'}
                                value={this.state.invoice.customer}
                                onChange={this.handleChange}
                                labelPosition="above"
                            />
                            <Text
                                label="Contacto"
                                name={'contact'}
                                value={this.state.invoice.contact}
                                onChange={this.handleChange}
                                labelPosition="above"
                            />
                        </Col>
                    </Box>
                    <Box collapsable title={'Carroceria y Pintura'}>
                        <Col xs={12}>
                            <SaleTable2
                                details={this.state.invoice.details}
                                handleAddItem={this.handleAddItem}
                                //handleDelItem={handleDelItem}
                                //handleChange={handleChange}
                                listProducts={this.state.products.map(p=>{return {value:p.id, label:p.description}})}
                                formatCurrency={this.state.currency}
                            />
                        </Col>
                        <Col xs={12}>
                            <Button
                                icon={'fa-save'}
                                type={'danger'}
                                text={'Guardar'}
                                onClick={this.handleSubmit}
                                pullRight
                            />
                        </Col>
                    </Box>
                </Col>
            </Row>
        </Content>);
    }
}

export default InvoiceCreate;