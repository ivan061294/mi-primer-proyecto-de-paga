import React, { Component } from 'react';

import {Box, Col, Row, Button} from "adminlte-2-react";
import Moment from 'react-moment';
import 'moment/locale/es';
import Content from "../components/Content";
import SimpleTable from "../components/SimpleTable";
import {formatCurrency} from "../util/Common"

class Quote extends Component {
    ws = new WebSocket('ws://localhost:8080/ws');
    constructor(props) {
        super(props);
        this.state = {
            quotes: null
        };
        this.columns = [
            {title: '#', data: 'id'},
            {title: 'Fecha', data: 'issue', render: e => (
                <Moment title={e} fromNow>{e}</Moment>
            )},
            {title: 'Vendedor', data: 'seller'},
            {title: 'Cliente', data: 'customer'},
            {title: 'Tipo Doc.', data: 'doctype'},
            {title: '# Doc.', data: 'docnum'},
            {title: 'Contacto', data: 'contact'},
            {title: 'Total', data: 'total', render: (e, d) => 
                formatCurrency(d.total, d.currency)
            },
            {title: 'Estado', data: 'status'},
            {title: 'Accion', data: 'action', render: (e, d) => {
                const ordeEnable = d.status==='Aceptado'?true:false;
                const viewEnable = true;
                const editEnable = d.status==='Pendiente'?true:false;
                const dropEnable = d.status==='Pendiente'?true:false;
                return(
                    <div>
                        <Button disabled={!ordeEnable} size={'xs'} icon={'fa-truck'} onClick={()=>this.props.history.push('/order/create?quoteid='+e)} />
                        <Button disabled={!viewEnable} size={'xs'} icon={'fa-eye'} onClick={()=>this.props.history.push('/quote/' + e)} />
                        <Button disabled={!editEnable} size={'xs'} icon={'fa-edit'} onClick={()=>this.props.history.push('/quote/' + e + '/edit')} />
                        <Button disabled={!dropEnable} size={'xs'} icon={'fa-trash'} onClick={this.handleDeleteQuote.bind(this, e)} />
                    </div>)}}
        ];
        this.handleDeleteQuote = this.handleDeleteQuote.bind(this);
    }
    handleDeleteQuote(id) {
        if (window.confirm("Esta seguro de eliminar la cotizacion #" + id + "?")) {
            fetch('http://localhost:8080/api/v1/quotations/' + id, {
                method: 'DELETE',
            })
            .then(res => res.json())
            .then(data => {
                if (data.code > 0) {
                    alert(data.message);
                } else {
                    this.ws.send("delete quote");
                    //this.ws.close();
                    this.setState({
                        quotes: this.state.quotes.filter(quote => quote.id !== id)
                    });
                }
            })
            .catch(console.log)
        }
    }
    componentDidMount() {
        fetch('http://localhost:8080/api/v1/quotations')
            .then(res => res.json())
            .then(data => {
                data = data?data.map(x => Object.assign({}, x, { "action": x.id })):null;
                this.setState({quotes: data})
            })
            .catch(console.log)
    }
    componentWillUnmount() {
        if (this.ws)
            this.ws.close()
    }
    render() {
        const validRenderComplete = () => {
            return true
                && this.state.quotes
                && true;
        };
        return(<Content title={this.props.title} subTitle="Lista" loaded={validRenderComplete()}>
            <Box header={<Button
                to={'quote/create'}
                text={'Nueva Cotizacion'}
                type={'danger'}
                pullRight
            />}>
                <Row>
                    <Col xs={12}>
                        <div className={'table-responsive'}>
                            <SimpleTable
                                columns={this.columns}
                                data={this.state.quotes}
                                messageNoRecords={'No existe ninguna cotizacion'}
                            />
                        </div>
                    </Col>
                </Row>
            </Box>
        </Content>);
    }
}

export default Quote;