import React, { Component } from 'react';

import {Box, Col, Row} from "adminlte-2-react";
import Content from "../components/Content";
import SimpleTable from "../components/SimpleTable";

import {Button} from "adminlte-2-react";
import Moment from "react-moment";

class Order extends Component {
    ws = new WebSocket('ws://localhost:8080/ws');
    constructor(props) {
        super(props);
        this.state = {
            orders: null
        };
        this.columns = [
            {title: '#', data: 'id'},
            {title: 'Fecha', data: 'issue', render: e => (
                <Moment title={e} fromNow>{e}</Moment>
            )},
            {title: 'Vendedor', data: 'seller'},
            {title: 'Cliente', data: 'customer'},
            {title: 'Total H.H', data: 'totalhours'},
            {title: 'Total Paños', data: 'totalcloths'},
            {title: 'Fecha de inicio', data: 'startdate', render: e => (
                <Moment title={e} fromNow>{e}</Moment>
            )},
            {title: 'Fecha de entrega', data: 'enddate', render: e => (
                <Moment title={e} fromNow>{e}</Moment>
            )},
            {title: 'Estado', data: 'status'},
            {title: 'Accion', data: 'action', render: e => (
                    <div>
                        <Button size={'xs'} icon={'fa-certificate'} to={'/certify/create?order=' + e} />
                        <Button size={'xs'} icon={'fa-eye'} to={'/order/' + e} />
                        <Button size={'xs'} icon={'fa-edit'} to={'/order/' + e + '/edit'} />
                        <Button size={'xs'} icon={'fa-trash'} to={'#'} onClick={this.handleDeleteorder.bind(this, e)} />
                    </div>)}
        ];
        this.handleDeleteorder = this.handleDeleteorder.bind(this);
    }
    handleDeleteorder(id) {
        if (window.confirm("Esta seguro de eliminar la orden de trabajo #" + id + "?")) {
            fetch('http://localhost:8080/api/v1/orders/' + id, {
                method: 'DELETE',
            })
            .then(res => res.json())
            .then(data => {
                console.log(data);
                if (data.code > 0){
                    alert(data.message);
                } else {
                    this.ws.send("delete order");
                    this.setState({
                        orders: this.state.orders.filter(order => order.id !== id)
                    });
                }
            })
            .catch(console.log)
        }
    }
    componentDidMount() {
        fetch('http://localhost:8080/api/v1/orders')
            .then(res => res.json())
            .then(data => {
                data = data?data.map(x => Object.assign({}, x, { "action": x.id })):null
                this.setState({orders: data})
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
                && this.state.orders
                && true;
        };
        return(<Content title={this.props.title} subTitle="Lista" loaded={validRenderComplete()}>
            <Box header={<Button
                to={'order/create'}
                text={'Nueva Orden de trabajo'}
                type={'danger'}
                pullRight
            />}>
                <Row>
                    <Col xs={12}>
                        <div className={'table-responsive'}>
                            <SimpleTable
                                columns={this.columns}
                                data={this.state.orders}
                                messageNoRecords={'No existe ninguna orden de trabajo'}
                            />
                        </div>
                    </Col>
                </Row>
            </Box>
        </Content>);
    }
}

export default Order;