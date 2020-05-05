import React, { Component } from 'react';
import Content from "../../components/Content";
import {Box, Col, Row} from "adminlte-2-react";

class InvoiceView extends Component {
    constructor(props) {
        super(props);
        this.state = {
            quotations: {},
        }
        this.columns = [
            {title: 'Descripcion', width: 150, data: 'description'},
            {title: 'Cantidad', width: 10, align: 'right', data: 'amount'},
            {title: 'Precio unitario', width: 10, align: 'right', data: 'unitprice'},
            {title: 'Precio total', width: 10, align: 'right', data: null}
        ];
    }
    componentDidMount() {
        const id = this.props.match.params.id
        fetch('http://localhost:8080/api/v1/quotations/'+ id)
            .then(res => res.json())
            .then(data => {
                console.log(data)
                this.setState({quotations: data})
            })
            .catch(console.log);
    }
    render() {
        const id = this.props.match.params.id
        const title = "Factura";
        const breadCrumb = [
            {title: "Factura", link: "/invoice"},
            {title: id},
        ];
        console.log(this.state.quotations.total)
        return (<Content title={title} subTitle="Detalle" breadCrumb={breadCrumb} browserTitle={title}>
            <Box>
                <Row>
                    <Col md={3} />
                    <Col md={4}>
                        <h4>{this.state.quotations.customer}</h4>
                        <p>{this.state.quotations.doctype} {this.state.quotations.docnum}</p>
                        <p>{this.state.quotations.contact}</p>
                    </Col>
                    <Col md={2}>
                        <h4>Cotizacion #{this.state.quotations.id}</h4>
                        <p>{this.state.quotations.issue}</p>
                        <p>{this.state.quotations.status}</p>
                    </Col>
                    <Col md={3} />
                </Row>
                <Row>
                    <Col md={3} />
                    <Col md={6}>
                    </Col>
                    <Col md={3} />
                </Row>
            </Box>
        </Content>)
    }
}
export default InvoiceView;