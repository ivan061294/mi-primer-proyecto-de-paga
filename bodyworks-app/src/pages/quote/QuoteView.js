import React, { Component } from 'react';
import Content from "../../components/Content";
import {Box, Col, Row} from "adminlte-2-react";

import '../../components/CustomTable.css'

class QuoteView extends Component {
    constructor(props) {
        super(props);
        this.state = {
            quotations: {},
        };
        this.columns = [
            {title: 'Descripcion', width: 200, data: 'description', render: (e, idx, d) => (<div>
                    {d.description}<br /><i>{d.description2}</i>
            </div>)},
            {title: 'Cantidad', width: 10, align: 'right', data: 'amount'},
            {title: 'Precio unitario', width: 10, align: 'right', data: 'unitprice'},
            {title: 'Precio total', width: 10, align: 'right', data: null}
        ];
    }
    componentDidMount() {
        const id = this.props.match.params.id;
        fetch('http://localhost:8080/api/v1/viewquotes/'+id)
            .then(res => res.json())
            .then(async data => {
                this.setState({quotations: data})
            })
            .catch(console.log);
    }
    render() {
        const id = this.props.match.params.id
        const breadCrumb = [
            {title: "Cotizacion", link: "/quote"},
            {title: id},
        ];
        const details = this.state.quotations.detail;
        const formatCurrency = this.state.quotations.currency;

        const currency=formatCurrency?formatCurrency:'PEN';
        const formatter = new Intl.NumberFormat(currency==='PEN'?'es-PE':'en-US', {
            style: 'currency',
            currency: currency,
            minimumFractionDigits: 2
        });
        const convertCurrency=currency==='PEN'?1:1/localStorage.getItem('dollarPrice');

        const price = details?details.map(d=>d.amount * d.unitprice).reduce((a,b)=>a+b):null * convertCurrency;
        const igv = price * 0.18;
        const total = price + igv;
        const validRenderComplete = () => {
            return true
                && this.state.quotations.id>=0
                && true;
        };
        return (<Content title={this.props.title} subTitle="Detalle" breadCrumb={breadCrumb} loaded={validRenderComplete()}>
            <Box title={''} solid>
                <Row>
                    <Col xs={1}/>
                    <Col xs={5}>
                        <h3>BODYWORKS PERÚ S.A.C.</h3>
                        <h4>RUC 20549734899</h4>
                        <h5>CALLE TANGANICA 120 LA MOLINA</h5>
                        <h5>TELF: 7152557</h5>
                    </Col>
                    <Col xs={5}>
                        <h3>
                            <img src={'../../logo.png'} height={64} align={'right'} />
                        </h3>
                    </Col>
                    <Col xs={1}/>
                </Row>
                <hr/>
                <Row>
                    <Col xs={1} />
                    <Col xs={7}>
                        <table className={'table table-bordered-bw'}>
                            <tbody>
                                <tr>
                                    <th style={{ width: 120 }}>Cliente: </th>
                                    <td>{this.state.quotations.customer}</td>
                                </tr>
                                <tr>
                                    <th>Atención: </th>
                                    <td>{this.state.quotations.contact}</td>
                                </tr>
                                <tr>
                                    <th>Color: </th>
                                    <td>{this.state.quotations.color}</td>
                                </tr>
                                <tr>
                                    <th>Elaborado por: </th>
                                    <td>{this.state.quotations.seller}</td>
                                </tr>
                            </tbody>
                        </table>
                    </Col>
                    <Col xs={3}>
                        <table className={'table table-bordered-bw'}>
                            <tbody>
                            <tr>
                                <th style={{ width: 120 }}>Marca</th>
                                <td>{this.state.quotations.brand}</td>
                            </tr>
                            <tr>
                                <th>Modelo</th>
                                <td>{this.state.quotations.model}</td>
                            </tr>
                            <tr>
                                <th>Placa</th>
                                <td>{this.state.quotations.plate}</td>
                            </tr>
                            <tr>
                                <th>Serie</th>
                                <td>{this.state.quotations.serie}</td>
                            </tr>
                            </tbody>
                        </table>
                    </Col>
                    <Col xs={1} />
                </Row>
                <Row>
                    <Col xs={1}/>
                    <Col xs={10}>
                        <table className={'table table-bordered-bw'}>
                            <thead>
                            <tr>
                                <th className={'text-center'} style={{ width: 40 }}>Item</th>
                                <th className={'text-center'} style={{ width: 300 }}>Descripcion</th>
                                <th className={'text-center'} style={{ width: 60 }}>Cantidad</th>
                                <th className={'text-center'} style={{ width: 60 }}>P. Unit.</th>
                                <th className={'text-center'} style={{ width: 60 }}>Total</th>
                            </tr>
                            </thead>
                            <tbody>
                                {details?details.map((detail, idx)=>{
                                    return(<tr key={idx}>
                                        <td className={'text-center'}>{idx + 1}</td>
                                        <td className={'text-left'}>{detail.product}<br /><i>{detail.description}</i></td>
                                        <td className={'text-center'}>{detail.amount}</td>
                                        <td className={'text-right'}>{formatter.format(detail.unitprice * convertCurrency)}</td>
                                        <td className={'text-right'}>{formatter.format(detail.amount * detail.unitprice * convertCurrency)}</td>
                                    </tr>)
                                }):null}
                                <tr>
                                    <td colSpan={3} rowSpan={3} className={'no-border'}></td>
                                    <th>Valor Venta</th>
                                    <td className={'text-right'}>{formatter.format(price)}</td>
                                </tr>
                                <tr>
                                    <th>IGV</th>
                                    <td className={'text-right'}>{formatter.format(igv)}</td>
                                </tr>
                                <tr>
                                    <th>Precio Total</th>
                                    <td className={'text-right'}>{formatter.format(total)}</td>
                                </tr>
                            </tbody>
                        </table>
                    </Col>
                    <Col xs={1}/>
                </Row>
            </Box>
        </Content>)
    }
}
export default QuoteView;