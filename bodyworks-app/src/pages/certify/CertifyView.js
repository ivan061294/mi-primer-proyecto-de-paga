import React, { Component } from 'react';
import {Box, Row, Col} from "adminlte-2-react";
import Content from "../../components/Content";
import Moment from "react-moment";

import '../../components/CustomTable.css'

class CertifyView extends Component {
    constructor(props) {
        super(props)
        this.state = {
            certy: null
        }
    }
    componentDidMount() {
        const id = this.props.match.params.id;
        fetch('http://localhost:8080/api/v1/viewcertifys/'+id)
            .then(res => res.json())
            .then(async data => {
                this.setState({order: data})
            })
            .catch(console.log);
    }
    render() {
        const id = this.props.match.params.id;
        const breadCrumb = [
            {title: "Acta de entrega", link: "/certify"},
            {title: id},
        ];
        const validRenderComplete = () => {
            return this.state.order &&
                this.state.employees.length > 0;
        };
        return (<Content title={this.props.title} subTitle="Detalle" breadCrumb={breadCrumb} loaded={validRenderComplete()}>
            {validRenderComplete()?<Box title={''}>
                <Row>
                </Row>
            </Box>:null}
        </Content>)
    }
}
export default CertifyView;