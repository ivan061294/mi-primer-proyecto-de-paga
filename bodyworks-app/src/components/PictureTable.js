import React from 'react'
import {Box, Inputs} from "adminlte-2-react";
const {Text} = Inputs;
class PictureTable extends React.Component {
    render() {
        const {title} = this.props;
        return(<Box collapsable title={title}>
            <table className={'table'}>
                <thead>
                    <tr>
                        <th>Color</th>
                        <th>Tipo</th>
                        <th>Cantidad</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td><Text name={'pieces'} labelPosition={'none'} /></td>
                        <td><Text name={'pieces'} labelPosition={'none'} /></td>
                        <td><Text name={'pieces'} labelPosition={'none'} /></td>
                    </tr>
                </tbody>
            </table>
        </Box>)
    }
}
export default PictureTable;