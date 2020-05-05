import React, { Component } from 'react';

import {Box, Col, Row} from "adminlte-2-react";

import {Button} from "adminlte-2-react";

import Content from "../components/Content";
import SimpleTable from "../components/SimpleTable";
import Moment from "react-moment";
import {formatCurrency} from "../util/Common"

class Invoice extends Component {
    constructor(props) {
        super(props);
        this.state = {
            invoices: null
        };
        this.columns = [
            {title: '#', data: 'id'},
            {title: 'Fecha', data: 'issue', render: e => (
                    <Moment title={e} fromNow>{e}</Moment>
                )},
            {title: 'Cliente', data: 'customer'},
            {title: 'Tipo Doc.', data: 'doctype'},
            {title: '# Doc.', data: 'docnum'},
            {title: 'Total', data: 'total', render: (e, d) =>
                formatCurrency(d.total, d.currency)
            },
            {title: 'Estado', data: 'status'},
            {title: 'Observaciones', data: 'observation'},
            {title: 'Archivo firmado', data: 'xmlsign', render: e =>
                <a href={'http://localhost:8080/api/v1/documents/'+e}>{e}</a>
            },
            {title: 'Archivo sunat', data: 'xmlsunat', render: e =>
                <a href={'http://localhost:8080/api/v1/documents/'+e}>{e}</a>
            },
            {title: 'Accion', data: 'action', render: e => (
                <>
                    <Button size={'xs'} icon={'fa-eye'} to={'/invoice/' + e} />
                    <Button size={'xs'} icon={'fa-edit'} to={'/invoice/' + e + '/edit'} />
                    <Button size={'xs'} icon={'fa-trash'} to={'#'} onClick={this.handleDeleteInvoice.bind(this, e)} />
                </>)
            }
        ];
        this.handleDeleteInvoice = this.handleDeleteInvoice.bind(this);
    }
    handleDeleteInvoice(id) {
        if (window.confirm("Esta seguro de eliminar la Factura #" + id + "?")) {
            fetch('http://localhost:8080/api/v1/invoices/' + id, {
                method: 'DELETE',
            })
            .then(res => res.json())
            .then(data => {
                if (data.code > 0)
                    alert(data.message)
                else
                    var invoices = this.state.invoices.filter(invoice => invoice.id !== id)
                    this.setState({invoices: invoices})
            })
            .catch(console.log)
        }
    }
    componentDidMount() {
        fetch('http://localhost:8080/api/v1/invoices')
            .then(res => res.json())
            .then(data => {
                console.log(data)
                data = data?data.map(x => Object.assign({}, x, { "action": x.id })):null
                this.setState({invoices: data})
            })
            .catch(console.log)
    }
    render() {
        const validRenderComplete = () => {
            return true
                && this.state.invoices
                && true;
        };
        return(<Content title={this.props.title} subTitle="Lista" loaded={validRenderComplete()}>
            <Box header={<Button
                to={'invoice/create'}
                text={'Nueva Factura'}
                type={'danger'}
                pullRight
            />}>
                <Row>
                    <Col xs={12}>
                        <div className={'table-responsive'}>
                            <SimpleTable
                                columns={this.columns}
                                data={this.state.invoices}
                                messageNoRecords={'No existe ninguna factura'}
                            />
                        </div>
                    </Col>
                </Row>
            </Box>
        </Content>);
    }
}

export default Invoice;