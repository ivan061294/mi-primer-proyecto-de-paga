import React, { Component } from 'react';
import {Box, Button, Row, Col, Inputs} from "adminlte-2-react";
import Content from "../../components/Content";

const { Text } = Inputs;

class Setting extends Component {
    constructor(props) {
        super(props);
        this.state = {
            dollar: localStorage.getItem('dollarPrice')
        }
        this.handleSubmit = this.handleSubmit.bind(this);
    }
    handleSubmit() {
        if (window.confirm("Esta seguro?")) {
            fetch('http://localhost:8080/api/v1/settings', {
                method: 'PUT',
                body: JSON.stringify({
                    name: 'DOLLAR_PRICE',
                    value: this.state.dollar
                })
            })
            .then(res => res.json())
            .then(data => {
                if (data.code > 0)
                    alert(data.message)
                else
                    localStorage.setItem('dollarPrice', this.state.dollar);
            })
            .catch(console.log)
        }
    }
    render() {
        const handleChange = e => {
            const { name, value } = e.target;
            this.setState({ [name]: value });
        };
        const validRenderComplete = () => {
            return true;
        };
        return (<Content title={this.props.title} subTitle="General" loaded={validRenderComplete()}>
            <Box header={<Button
                //to={'order/create'}
                onClick={this.handleSubmit}
                text={'Guardar'}
                type={'danger'}
                pullRight
            />}>
                <Row>
                    <Col xs={12}>
                        <Text
                            name={'dollar'}
                            labelPosition={'above'}
                            label={'Precio del dolar'}
                            value={this.state.dollar}
                            onChange={handleChange}
                        />
                    </Col>
                </Row>
            </Box>
        </Content>);
    }
}
export default Setting;