import React from 'react'
import {Box, Inputs} from "adminlte-2-react";
import './CustomTable.css'
const {Text, Checkbox} = Inputs;
class PictureTable extends React.Component {
    render() {
        const {suppliers, defaults, handleChange} = this.props;
        return(<Box collapsable title={'Insumos'}>
            <table className={'table'}>
                <thead>
                    <tr>
                        <th>Tipo</th>
                        <th>Unidad</th>
                        <th>Cantidad</th>
                    </tr>
                </thead>
                <tbody>
                    {suppliers.map(supplier => {
                        /*defaults.map((d)=>{
                            if (supplier.id===d.product && !supplier.amount)
                                supplier.amount=d.amount;
                        })*/
                        return(
                        <tr>
                            <td>{supplier.description}</td>
                            <td>{supplier.measurement}</td>
                            <td className={'border-clean'}>
                                <Text
                                    value={supplier.amount}
                                    inputType={'number'}
                                    size={'xs'}
                                    name={supplier.id}
                                    labelPosition={'none'}
                                    onChange={handleChange}
                                />
                            </td>
                        </tr>
                    )})}
                </tbody>
            </table>
        </Box>)
    }
}
export default PictureTable;