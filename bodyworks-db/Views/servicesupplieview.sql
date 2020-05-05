CREATE VIEW quotesupplieview AS
SELECT quotedetails.quotation_id,
       servicesupplie.supplie_id,
       (quotedetails.amount * servicesupplie.relation) AS autoamount
FROM quotedetails
JOIN servicesupplie
  ON quotedetails.product_id = servicesupplie.service_id;