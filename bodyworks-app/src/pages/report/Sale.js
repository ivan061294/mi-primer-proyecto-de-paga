import React, { Component } from 'react';
import {Box, Col, Infobox, Infobox2, Row} from "adminlte-2-react";
import SaleChart from '../../components/SaleChart'
import SaleLine from '../../components/SaleLine'
import Content from "../../components/Content";
class Sale extends Component {
    constructor(props) {
        super(props);
        this.state = {
            quoteCount: 0,
            orderCount: 0,
            certifyCount: 0,
            invoiceCount: 0,
            saleForMonth: null,
        };
    }
    ws = new WebSocket('ws://localhost:8080/ws');
    componentDidMount() {
        this.ws.onopen = () => {
            console.log('open');
            this.ws.send('init report')
        };
        this.ws.onmessage = e => {
            let report = JSON.parse(e.data);
            this.setState({
                quoteCount: report.countQuote,
                orderCount: report.countOrder,
                certifyCount: report.countCertify,
                invoiceCount: report.countInvoice,
                saleForMonth: report.saleForMonth
            });
        };
        this.ws.onclose = () => {
            console.log('close');
        };
    }
    componentWillUnmount() {
        this.ws.close()
    }
    render() {
        const validRenderComplete = () => {
            return true
                && this.state.saleForMonth
        };
        return (<Content title={this.props.title} subTitle="Dashboard" loaded={validRenderComplete()}>
            <Box>
                <Row>
                    <Col md={3}>
                        <a href={'#'} onClick={()=>this.props.history.push('/quote')} >
                            <Infobox
                                content={''}
                                color={'yellow'}
                                icon={'fa-tag'}
                                number={this.state.quoteCount}
                                text={'Cotizaciones'}
                            />
                        </a>
                    </Col>
                    <Col md={3}>
                        <a href={'#'} onClick={()=>this.props.history.push('/order')} >
                            <Infobox
                                content={''}
                                color={'light-blue'}
                                icon={'fa-tag'}
                                number={this.state.orderCount}
                                text={'Ordenes de trabajo'}
                            />
                        </a>
                    </Col>
                    <Col md={3}>
                        <a href={'#'} onClick={()=>this.props.history.push('/certify')} >
                            <Infobox
                                content={''}
                                color={'red'}
                                icon={'fa-tag'}
                                number={this.state.certifyCount}
                                text={'Actas de entrega'}
                            />
                        </a>
                    </Col>
                    <Col md={3}>
                        <a href={'#'} onClick={()=>this.props.history.push('/invoice')} >
                            <Infobox
                                content={''}
                                color={'green'}
                                icon={'fa-tag'}
                                number={this.state.invoiceCount}
                                text={'Facturas'}
                            />
                        </a>
                    </Col>
                </Row>
                <Row>
                    <Col md={6}>
                        <SaleChart
                            countQuote={this.state.quoteCount}
                            countOrder={this.state.orderCount}
                            countCertify={this.state.certifyCount}
                            countInvoice={this.state.invoiceCount}
                        />
                    </Col>
                    <Col md={6}>
                        <SaleLine saleForMonth={this.state.saleForMonth} />
                    </Col>
                </Row>
            </Box>
        </Content>);
    }
}
export default Sale;