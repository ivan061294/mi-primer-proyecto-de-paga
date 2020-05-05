import React, {Component} from 'react';

import {Box, Row, Col} from "adminlte-2-react";
import Content from "../../components/Content";
import SaleTable2 from "../../components/SaleTable2";
const uuid = require('uuid/v4');

class InvoiceEdit extends Component {
    constructor(props) {
        super(props);
        this.state = {
            products : [],
            details : [],
            currency: 'PEN'
        }
    }
    componentDidMount() {
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
        const listProducts = this.state.products;
        const handleChange = e => {
            let id = e.target.name.split('-')[1];
            let name = e.target.name.split('-')[0];
            let value = e.target.value;
            let details = [...this.state.details];
            if (id && name) {
                details[id][name]=value;
            }
            if (name === 'product') {
                details[id]["unitprice"]=listProducts.find(p=>p.id===parseInt(value)).unitprice;
            }
            details=details.map(detail => {
                detail.price = detail.quantity * detail.unitprice;
                return detail
            });
            this.setState({
                details: details,
                currency: e.target.name==='currency'?e.target.value:this.state.currency
            })
        };
        const handleAddItem = e => {
            let detail = new Object();
            detail.id = uuid();
            detail.quantity = 1;
            detail.product = 0;
            detail.description = '';
            detail.unitprice = 0.00;
            detail.price = 0.00;
            this.setState((preview) => ({
                details: [...preview.details, detail]
            }));
        };
        const handleDelItem = id => {
            this.setState({
                details: this.state.details.filter(detail=>detail.id !== id)
            })
        };
        console.log(this.state.products);
        return (<Content title={'Factura'} subTitle={'Modificación'} breadCrumb={[
                {title: "Factura", link: "/invoice"},
                {title: this.props.match.params.id}
            ]}>
            <Box>
                <Row>
                    <Col md={12}>
                        {this.state.products.length>0?
                        <SaleTable2
                            details={this.state.details}
                            handleAddItem={handleAddItem}
                            handleDelItem={handleDelItem}
                            handleChange={handleChange}
                            listProducts={this.state.products.map(p=>{return {value:p.id, label:p.description}})}
                            formatCurrency={this.state.currency}
                        />:null}
                    </Col>
                </Row>
            </Box>
        </Content>);
    }
}

export default InvoiceEdit;