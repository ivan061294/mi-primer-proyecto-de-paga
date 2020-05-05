import React, { Component } from 'react';
import {Box, Row, Col} from "adminlte-2-react";
import Content from "../../components/Content";
import Moment from "react-moment";

import '../../components/CustomTable.css'

class OrderView extends Component {
    constructor(props) {
        super(props)
        this.state = {
            order: null,
            employees: [],
        }
    }
    componentDidMount() {
        const id = this.props.match.params.id;
        fetch('http://localhost:8080/api/v1/vieworders/'+id)
            .then(res => res.json())
            .then(async data => {
                this.setState({order: data})
            })
            .catch(console.log);
        fetch('http://localhost:8080/api/v1/employees')
            .then(res => res.json())
            .then(data => {this.setState({employees: data})})
            .catch(console.log);
    }
    render() {
        const id = this.props.match.params.id;
        const breadCrumb = [
            {title: "Orden de trabajo", link: "/order"},
            {title: id},
        ];
        const validRenderComplete = () => {
            return this.state.order &&
                this.state.employees.length > 0;
        };
        const getSupervidors = (employees) => {
            return this.state.employees
                .filter(e=>e.care==="SUPERVISOR")
                .filter(e=>employees.includes(e.id))
                .map(e=>e.name + " " + e.lastname)
                .join(", ");
        };
        const getColaboradors = (employees) => {
            return this.state.employees
                .filter(e=>e.care==="COLABORADOR")
                .filter(e=>employees.includes(e.id))
                .map(e=>e.name + " " + e.lastname)
                .join(", ");
        };
        let employees;
        if (validRenderComplete()) {
            employees = Array.prototype.concat.apply([], this.state.order.detail.map(d=>d.subdetail)).map(d=>d.employee);
        }
        return (<Content title={this.props.title} subTitle="Detalle" breadCrumb={breadCrumb} loaded={validRenderComplete()}>
            {validRenderComplete()?<Box title={''}>
                <Row>
                    <Col md={1} />
                    <Col md={10} >
                        <Row>
                            <Col md={5}>
                                <h3>
                                    <img src={'../../logo.png'} height={64} align={'left'} />
                                </h3>
                            </Col>
                        </Row>
                        <Row>
                            <Col md={12}>
                                <h3>Datos Generales</h3>
                                <table className={'table table-bordered-bw'}>
                                    <tbody>
                                        <tr>
                                            <th>Cliente: </th>
                                            <td colSpan={5}>{this.state.order.customer}</td>
                                            <th>Fecha: </th>
                                            <td>{this.state.order.issue}</td>
                                        </tr>
                                        <tr>
                                            <th>Marca: </th>
                                            <td>{this.state.order.brand}</td>
                                            <th>Modelo: </th>
                                            <td>{this.state.order.model}</td>
                                            <th>Color: </th>
                                            <td>{this.state.order.color}</td>
                                            <th>Placa/Chasis: </th>
                                            <td>{this.state.order.plate}</td>
                                        </tr>
                                    </tbody>
                                </table>
                                <h3>Tipo de trabajo</h3>
                                <table className={'table table-bordered-bw'}>
                                    <tbody>
                                        <tr>
                                            <th style={{ width: 80 }}>PDI</th>
                                            <td className={'text-center'} style={{ width: 30 }}>{this.state.order.worktype==='PDI'?'X':null}</td>
                                            <th style={{ width: 80 }}>SINIESTRO</th>
                                            <td className={'text-center'} style={{ width: 30 }}>{this.state.order.worktype==='SINIESTRO'?'X':null}</td>
                                            <th style={{ width: 80 }}>UNIDAD COMPLETA</th>
                                            <td className={'text-center'} style={{ width: 30 }}>{this.state.order.worktype==='UNIDAD COMPLETA'?'X':null}</td>
                                            <th style={{ width: 80 }}>PIEZAS</th>
                                            <td className={'text-center'} style={{ width: 30 }}>{this.state.order.worktype==='PIEZAS'?'X':null}</td>
                                        </tr>
                                    </tbody>
                                </table>
                                <h3>Tiempos y horas hombres</h3>
                                <table className={'table table-bordered-bw'}>
                                    <tbody>
                                        <tr>
                                            <th>Fecha de aprobacion:</th>
                                            <td>
                                                <Moment format="DD/MM/YYYY">{this.state.order.issue}</Moment>
                                            </td>
                                            <th>Fecha de inicio:</th>
                                            <td>
                                                <Moment format="DD/MM/YYYY">{this.state.order.startdate}</Moment>
                                            </td>
                                            <th>Fecha de entrega:</th>
                                            <td>
                                                <Moment format="DD/MM/YYYY">{this.state.order.enddate}</Moment>
                                            </td>
                                            <th>N° Colaboradores:</th>
                                            <td className={'text-center'}>{employees.length}</td>
                                        </tr>
                                        <tr>
                                            <th>Encargado del trabajo:</th>
                                            <td colSpan={5}>{getSupervidors(employees)}</td>
                                            <th>Horas hombre (H.H.):</th>
                                            <td className={'text-center'}>{this.state.order.totalhours}</td>
                                        </tr>
                                        <tr>
                                            <th>Personal de apoyo:</th>
                                            <td colSpan={5}>{getColaboradors(employees)}</td>
                                            <th>Total:</th>
                                            <td className={'text-center'}>{this.state.order.totalhours}</td>
                                        </tr>
                                    </tbody>
                                </table>
                                <h3>Descripcion del trabajo</h3>
                                <table className={'table table-bordered-bw'}>
                                    <tbody>
                                    <tr>
                                        <th className={'text-center'} style={{ width: 30 }}>Item</th>
                                        <th className={'text-center'} style={{ width: 320 }}>Descripcion del trabajo</th>
                                        <th className={'text-center'} style={{ width: 30 }}>N° H.H</th>
                                        <th className={'text-center'} style={{ width: 30 }}>N° Paños</th>
                                    </tr>
                                    {this.state.order.detail?this.state.order.detail.map((detail, idx)=><tr key={idx}>
                                        <td className={'text-center'}>{idx + 1}</td>
                                        <td>{detail.description}</td>
                                        <td className={'text-center'}>{detail.workhours}</td>
                                        <td className={'text-center'}>{detail.cloths}</td>
                                    </tr>):null}
                                    <tr>
                                        <td colSpan={2} className={'text-right no-border'}><b>Total</b></td>
                                        <td className={'text-center'}>{this.state.order.detail?this.state.order.detail.map((e)=>e.workhours).reduce((a,b)=>a+b):0}</td>
                                        <td className={'text-center'}>{this.state.order.detail?this.state.order.detail.map((e)=>e.workhours).reduce((a,b)=>a+b):0}</td>
                                    </tr>
                                    </tbody>
                                </table>
                                <hr/>
                            </Col>
                        </Row>
                    </Col>
                </Row>
            </Box>:null}
        </Content>)
    }
}
export default OrderView;