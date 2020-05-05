import React, { Component } from 'react';
import {Button, Inputs, Row, Col} from "adminlte-2-react";

import {formatCurrency, priceCurrency} from "../util/Common"

const { Text, Select2 } = Inputs;

const uuid = require('uuid/v4');

class SaleTable2 extends Component {
    render(){
        const {
            details, handleAddItem, handleDelItem, handleChange, listProducts, currency, handleChangeCurrency
        } = this.props;
        //const currency=formatCurrency?formatCurrency:'PEN'
        /*const formatter = new Intl.NumberFormat(currency==='PEN'?'es-PE':'en-US', {
            style: 'currency',
            currency: currency,
            minimumFractionDigits: 2
        });*/
        //const convertCurrency=parseFloat(currency==='PEN'?1:1/localStorage.getItem('dollarPrice')).toFixed(2);
        /*const convertCurrency=currency==='PEN'?1:1/localStorage.getItem('dollarPrice')*/
        const thead = <tr>
            <th style={{ width: 80 }} className={'text-left'}>Accion</th>
            <th style={{ width: 300 }} className={'text-left'}>Servicio</th>
            <th style={{ width: 400 }} className={'text-left'}>Description</th>
            <th style={{ width: 150 }} className={'text-left'}>Cantidad</th>
            <th style={{ width: 150 }} className={'text-right'}>Precio Unitario</th>
            <th style={{ width: 150 }} className={'text-right'}>Precio total</th>
        </tr>;
        const tbody = details?details.map((detail, idx)=>{
            detail.idx = detail.idx ? detail.idx : uuid();
            return(
                <tr key={idx}>
                    <td>
                        <Button
                            icon={'fa-trash'}
                            type={'danger'}
                            onClick={()=>handleDelItem(detail.idx)}
                        />
                    </td>
                    <td>
                        <Select2
                            name={'product-' + idx}
                            value={detail.product}
                            options={listProducts}
                            labelPosition={'none'}
                            onChange={handleChange}
                        />
                    </td>
                    <td>
                        <Text
                            name={'description-' + idx}
                            value={detail.description}
                            labelPosition={'none'}
                            onChange={handleChange}
                        />
                    </td>
                    <td>
                        <Text
                            name={'quantity-' + idx}
                            value={detail.quantity}
                            inputType={'number'}
                            labelPosition={'none'}
                            onChange={handleChange}
                        />
                    </td>
                    <td>
                        <Text
                            name={'unitprice-' + idx}
                            //value={detail.unitprice * convertCurrency}
                            value={priceCurrency(detail.unitprice, currency)}
                            inputType={'number'}
                            labelPosition={'none'}
                            onChange={handleChange}
                        />
                    </td>
                    <td className={'text-right'}>
                        {
                    //formatter.format((detail.quantity*detail.unitprice)*convertCurrency)
                    formatCurrency(detail.quantity*detail.unitprice, currency)
                        }
                    </td>
                </tr>
            )
        }):null;
        //const totalPrice = details ? details.reduce((a,b)=>a + (b.quantity*b.unitprice), 0) * convertCurrency : 0;
        const totalPrice = details ? details.reduce((a,b)=>a + (b.quantity*b.unitprice), 0) : 0;
        const igv = totalPrice * 0.18;
        const total = totalPrice + igv;
        return(<div>
            <Row>
                <Col xs={2}>
                    <Select2
                        name={'currency'}
                        options={[{value:'PEN', label:'SOLES'},{value:'USD', label:'DOLARES'}]}
                        value={currency}
                        labelPosition={'none'}
                        onChange={handleChange}
                    />
                </Col>
                <Col xs={8} />
                <Col xs={2}>
                    <Button
                        icon={'fa-plus'}
                        text={'Agregar'}
                        type={'danger'}
                        onClick={handleAddItem}
                        pullRight
                    />
                </Col>
            </Row>
            <div className={'table-responsive'}>
                <table className={'table'}>
                    <thead>
                        {thead}
                    </thead>
                    <tbody>
                        {tbody}
                        <tr>
                            <th rowSpan={3} colSpan={4}/>
                            <th>Valor de venta</th>
                            <td className={'text-right'}>{
                                formatCurrency(totalPrice, currency)
                            //formatter.format(totalPrice)
                            }</td>
                        </tr>
                        <tr>
                            <th>IGV</th>
                            <td className={'text-right'}>{
                                formatCurrency(igv, currency)
                                //formatter.format(igv)
                            }</td>
                        </tr>
                        <tr>
                            <th>Total</th>
                            <td className={'text-right'}>{
                                formatCurrency(total, currency)
                                //formatter.format(total)
                            }</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>);
    }
}

export default SaleTable2;