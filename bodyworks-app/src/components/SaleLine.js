import React from 'react';
import {Line} from 'react-chartjs-2';
import moment from 'moment';

class SaleLine extends React.Component {
    render() {
        const {saleForMonth} = this.props;
        const data = {
            labels: saleForMonth.map(m => {
                let strmonth;
                strmonth = moment(m.month).subtract(0, "month").format('MMMM');
                return strmonth.charAt(0).toUpperCase() + strmonth.slice(1)
            }),
            datasets: [
                {
                    label: 'Ventas por mes',
                    fill: false,
                    lineTension: 0.1,
                    backgroundColor: 'rgba(75,192,192,0.4)',
                    borderColor: 'rgba(75,192,192,1)',
                    borderCapStyle: 'butt',
                    borderDash: [],
                    borderDashOffset: 0.0,
                    borderJoinStyle: 'miter',
                    pointBorderColor: 'rgba(75,192,192,1)',
                    pointBackgroundColor: '#fff',
                    pointBorderWidth: 1,
                    pointHoverRadius: 5,
                    pointHoverBackgroundColor: 'rgba(75,192,192,1)',
                    pointHoverBorderColor: 'rgba(220,220,220,1)',
                    pointHoverBorderWidth: 2,
                    pointRadius: 1,
                    pointHitRadius: 10,
                    data: saleForMonth.map(m=>moment(m.sale))
                }
            ]
        };
        return (
            <div>
                <Line data={data}/>
            </div>
        );
    }
}
export default SaleLine;