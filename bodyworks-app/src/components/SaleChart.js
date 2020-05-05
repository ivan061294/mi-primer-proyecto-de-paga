import React from 'react'
import {Pie} from 'react-chartjs-2';
class SaleChart extends React.Component {
    render() {
        const {countQuote, countOrder, countCertify, countInvoice} = this.props;
        const data = {
            labels: [
                'Cotizaciones',
                'Ordenes de trabajo',
                'Actas de entrega',
                'Facturas'
            ],
            datasets: [{
                data: [countQuote, countOrder, countCertify, countInvoice],
                backgroundColor: [
                    '#F39C12',
                    '#3C8DBC',
                    '#DD4B39',
                    '#00A65A'
                ],
                hoverBackgroundColor: [
                    '#F39C12',
                    '#3C8DBC',
                    '#DD4B39',
                    '#00A65A'
                ]
            }]
        };
        return (
            <div>
                <Pie data={data} />
            </div>
        );
    }
}
export default SaleChart;