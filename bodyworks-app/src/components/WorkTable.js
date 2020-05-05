import React, { Component } from 'react';
import {Button, Inputs, Row, Col} from "adminlte-2-react";

const { Text, Select2 } = Inputs;

const uuid = require('uuid/v4');

class WorkTable extends Component {
    render(){
        const {
            details, handleAddItem, handleDelItem, handleAddSubitem, handleDelSubitem, handleChange, employee
        } = this.props;
        const thead = <tr>
            <th className={'text-left'}>Accion</th>
            <th className={'text-left'}>Descripcion del trabajo</th>
            <th></th>
            <th className={'text-center'}># H.H.</th>
            <th className={'text-center'}># Paños</th>
        </tr>;
        const tbody = details?details.map((detail, idx)=>{
            detail.idx = detail.idx ? detail.idx : uuid();
            return(<>
                <tr key={idx}>
                    <td style={{ width: 60 }}>
                        <Button
                            icon={'fa-trash'}
                            type={'danger'}
                            onClick={()=>handleDelItem(detail.idx)}
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
                        <Button
                            text={'(+) Agregar detalle'}
                            onClick={()=>handleAddSubitem(detail.idx)}
                        />
                    </td>
                    <td style={{ width: 80 }} className={'text-center'}>{detail.workhours}</td>
                    <td style={{ width: 80 }} className={'text-center'}>{detail.cloths}</td>
                </tr>
                {detail.subdetail?detail.subdetail.map((subdetail, subidx)=>
                    <tr key={subidx}>
                        <td></td>
                        <td>
                            <Button
                                icon={'fa-trash'}
                                type={'danger'}
                                onClick={()=>handleDelSubitem(detail.idx, subdetail.idx)}
                                pullRight
                            />
                        </td>
                        <td>
                            <Select2
                                name={'employee-' + idx + '-' + subidx}
                                labelPosition={'none'}
                                options={employee}
                                value={subdetail.employee}
                                onChange={handleChange}
                            />
                        </td>
                        <td style={{ width: 80 }}>
                            <Text
                                name={'subworkhours-' + idx + '-' + subidx}
                                value={subdetail.workhours}
                                inputType={'number'}
                                labelPosition={'none'}
                                onChange={handleChange}
                            />
                        </td>
                        <td style={{ width: 80 }}>
                            <Text
                                name={'subcloths-' + idx + '-' + subidx}
                                value={subdetail.cloths}
                                inputType={'number'}
                                labelPosition={'none'}
                                onChange={handleChange}
                            />
                        </td>
                    </tr>
                ):null}
                </>)
        }):null;
        const totalworkhours = details.length>0?details.map(d=>d.workhours).reduce((a,b)=>a+b):0;
        const totalcloths = details.length>0?details.map(d=>d.cloths).reduce((a,b)=>a+b):0;
        return(<div>
            <Row>
                <Col xs={12}>
                    <Button
                        icon={'fa-plus'}
                        text={'Trabajo'}
                        type={'danger'}
                        onClick={handleAddItem}
                        pullRight
                    />
                </Col>
            </Row>
            <Row>
                <Col md={12}>
                    <div className={'table-responsive'}>
                        <table className={'table'}>
                            <thead>
                                {thead}
                            </thead>
                            <tbody>
                                {tbody}
                                <tr>
                                    <th colSpan={2}/>
                                    <th className={'text-right'}>Total</th>
                                    <td className={'text-center'}>{totalworkhours}</td>
                                    <td className={'text-center'}>{totalcloths}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </Col>
            </Row>
        </div>);
    }
}

export default WorkTable;