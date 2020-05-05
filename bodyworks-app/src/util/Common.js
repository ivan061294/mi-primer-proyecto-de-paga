export const formatCurrency = (total, currency='PEN') => {
    const locales = currency==='PEN'?'es-PE':'en-US';
    const formatter = new Intl.NumberFormat(locales, {
        style: 'currency',
        currency: currency,
        minimumFractionDigits: 2
    });
    return formatter.format(total * (currency==='PEN'?1:1/localStorage.getItem('dollarPrice')));
};

export const priceCurrency = (total, currency='PEN') => {
    return parseFloat(total * (currency==='PEN'?1:1/localStorage.getItem('dollarPrice'))).toFixed(2);
};