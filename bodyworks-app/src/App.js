import React, { Component } from 'react';
import AdminLTE, { Sidebar } from 'adminlte-2-react';
import Quote from './pages/Quote';
import QuoteView from './pages/quote/QuoteView';
import QuoteEdit from './pages/quote/QuoteEdit';
import QuoteCreate from './pages/quote/QuoteCreate';
import Order from './pages/Order';
import OrderView from './pages/order/OrderView';
import OrderEdit from './pages/order/OrderEdit';
import OrderCreate from './pages/order/OrderCreate';
import Certify from './pages/Certify';
import CertifyView from './pages/certify/CertifyView';
import CertifyEdit from './pages/certify/CertifyEdit';
import CertifyCreate from './pages/certify/CertifyCreate';
import Invoice from './pages/Invoice';
import InvoiceView from './pages/invoice/InvoiceView';
import InvoiceEdit from './pages/invoice/InvoiceEdit';
import InvoiceCreate from './pages/invoice/InvoiceCreate';
import ReportSale from './pages/report/Sale';
import Setting from './pages/setting/Setting';

import './App.css'

const { Item, Header } = Sidebar;

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            products: [],
            customers: []
        }
    }
    componentDidMount() {
        fetch('http://localhost:8080/api/v1/settings')
            .then(res => res.json())
            .then(data => {
                let dollarPrice = data.find(e=>e.name==="DOLLAR_PRICE");
                localStorage.setItem('dollarPrice', dollarPrice.value?dollarPrice.value:null);
            })
            .catch(console.log);
    }
    render() {
        const activeOnQuote=['/quote/create', '/quote/[0-9]', '/quote/[0-9]/edit'];
        const activeOnOrder=['/order/create', '/order/[0-9]', '/order/[0-9]/edit'];
        const activeOnCertify=['/certify/create', '/certify/[0-9]', '/certify/[0-9]/edit'];
        const activeOnInvoice=['/invoice/create', '/invoice/[0-9]', '/invoice/[0-9]/edit'];
        const titleQuote="Cotización";
        const titleOrder="Orden de trabajo";
        const titleInvoice="Factura";
        const titleCertify="Acta de entrega";
        const titleReportSale="Reporte de Ventas";
        const titleSetting="Configuracion";
        return (
            <AdminLTE title={["Body", "works"]} titleShort={["B", "W"]} theme="purple">
                <Sidebar.Core>
                    <Header text="NAVEGACION PRINCIPAL" />
                    <Item icon="fa-tag" key="sales" text="Ventas">
                        <Item text="Cotizacion" to="/quote" activeOn={activeOnQuote} />
                        <Item text="Orden de trabajo" to="/order" activeOn={activeOnOrder} />
                        <Item text="Acta de entrega" to="/certify" activeOn={activeOnCertify} />
                        <Item text="Factura" to="/invoice" activeOn={activeOnInvoice} />
                    </Item>
                    <Item icon="fa-shopping-cart" key="purchase" text="Compras">
                        <Item key="purchase" text="Compra" to="/purchase" />
                        <Item key="inventory" text="Inventario" to="/inventory" />
                    </Item>
                    <Item icon="fa-chart-pie" key="reports" text="Reportes">
                        <Item key="report-sale" text="Reporte de ventas" to="/report-sale" />
                        <Item key="report-pucharse" text="Reporte de compras" to="/report-pucharse" />
                    </Item>
                    <Item icon="fa-cog" key="setting" text="Configuracion">
                        <Item key="general" text="General" to="/setting" />
                    </Item>
                </Sidebar.Core>

                <Quote path={"/quote"} title={titleQuote} exact />
                <QuoteCreate path={"/quote/create"} title={titleQuote} exact />
                <QuoteView path={"/quote/:id"} title={titleQuote} exact />
                <QuoteEdit path={"/quote/:id/edit"} title={titleQuote} exact />

                <Order path={"/order"} title={titleOrder} exact />
                <OrderCreate path={"/order/create"} title={titleOrder} exact />
                <OrderCreate path={"/order/create?quote=:quoteid"} title={titleQuote} exact />
                <OrderView path={"/order/:id"} title={titleOrder} exact />
                <OrderEdit path={"/order/:id/edit"} title={titleOrder} exact />

                <Certify path={"/certify"} title={titleCertify} exact />
                <CertifyCreate path={"/certify/create"} title={titleCertify} exact />
                <CertifyView path={"/certify/:id"} title={titleCertify} exact />
                <CertifyEdit path={"/certify/:id/edit"} title={titleCertify} exact />

                <Invoice path={"/invoice"} title={titleInvoice} exact />
                <InvoiceCreate path={"/invoice/create"} title={titleInvoice} exact />
                <InvoiceView path={"/invoice/:id"} title={titleInvoice} exact />
                <InvoiceEdit path={"/invoice/:id/edit"} title={titleInvoice} exact />

                <ReportSale path={"/report-sale"} title={titleReportSale} />
                <Setting path={"/setting"} title={titleSetting} />

            </AdminLTE>
        );
    }
}

export default App;