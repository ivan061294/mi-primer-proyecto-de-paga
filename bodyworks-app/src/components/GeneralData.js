import React, { Component } from 'react';
import {Box, Col, Row, Inputs, Button} from "adminlte-2-react";
const {Text, Select2} = Inputs;
class GeneralData extends Component {
    constructor(props) {
        super(props)
    }
    render() {
        const {order, quotes, handleChange} = this.props;
        return(<Box collapsable title={'Datos generales'}>
            <Row>
                <Col md={3}>
                    <Select2
                        name={'quoteid'}
                        label={'Cotizacion'}
                        options={quotes.map(p=> p.id)}
                        value={order.quoteid}
                        labelPosition={'above'}
                        onChange={handleChange}
                    />
                </Col>
                <Col md={9}>
                    <Button
                        type={'danger'}
                        icon={'fa-save'}
                        text={'Guardar'}
                        onClick={this.handleSubmit}
                        pullRight
                    />
                </Col>
            </Row>
            <Row>
                <Col md={3}>
                    <Text
                        name={'brand'}
                        labelPosition={'above'}
                        label={'Marca'}
                        value={''}
                        onChange={handleChange}
                    />
                </Col>
                <Col md={3}>
                    <Text
                        name={'brand'}
                        labelPosition={'above'}
                        label={'Modelo'}
                        value={''}
                        onChange={handleChange}
                    />
                </Col>
                <Col md={3}>
                    <Text
                        name={'brand'}
                        labelPosition={'above'}
                        label={'Color'}
                        value={''}
                        onChange={handleChange}
                    />
                </Col>
                <Col md={3}>
                    <Text
                        name={'brand'}
                        labelPosition={'above'}
                        label={'Placa / Chasis'}
                        value={''}
                        onChange={handleChange}
                    />
                </Col>
            </Row>
        </Box>)
    }
}
export default GeneralData;