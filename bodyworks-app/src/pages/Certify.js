import React, { Component } from 'react';
import Content from "../components/Content";
import {Box, Button, Row, Col} from "adminlte-2-react";
import SimpleTable from "../components/SimpleTable";
import Moment from "react-moment";
class Certify extends Component {
    constructor(props)
    {
        super(props);
        this.state = {
            certifys: null
        };
        this.handleDeleteCertify = this.handleDeleteCertify.bind(this);
    }
    componentDidMount() {
        fetch('http://localhost:8080/api/v1/certifys')
            .then(res => res.json())
            .then(data => {
                data = data?data.map(x => Object.assign({}, x, { "action": x.id })):null
                this.setState({certifys: data})
            })
            .catch(console.log)
    }
    handleDeleteCertify(id) {
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
    render()
    {
        const columns = [
            {title: '#', data: 'id'},
            {title: 'Fecha', data: 'issue', render: e => (
                <Moment title={e} fromNow>{e}</Moment>
            )},
            {title: 'Cliente', data: 'customer'},
            {title: 'Contacto', data: 'contact'},
            {title: 'Descripcion', data: 'description'},
            {title: 'Accion', data: 'action', render: e => (
                <>
                    <Button size={'xs'} icon={'fa-money-check-alt'} to={'/invoice/' + e}  />
                    <Button size={'xs'} icon={'fa-eye'} to={'/certify/' + e} />
                    <Button size={'xs'} icon={'fa-edit'} to={'/certify/' + e + '/edit'} />
                    <Button size={'xs'} icon={'fa-trash'} to={'#'} onClick={this.handleDeleteCertify.bind(this, e)} />
                </>)}
        ];
        const validRenderComplete = () => {
            return this.state.certifys;
        };
        return(<Content title={this.props.title} subTitle={'Lista'} loaded={validRenderComplete()}>
            <Box header={<Button
                to={'order/create'}
                text={'Nueva Acta de entrega'}
                type={'danger'}
                pullRight
            />}>
                <Row>
                    <Col xs={12}>
                        <div className={'table-responsive'}>
                            <SimpleTable
                                columns={columns}
                                data={this.state.certifys}
                                messageNoRecords={'No existe ninguna Acta de entrega'}
                            />
                        </div>
                    </Col>
                </Row>
            </Box>
        </Content>);
    }
}
export default Certify;