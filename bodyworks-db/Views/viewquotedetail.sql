CREATE VIEW viewquotedetail AS
SELECT quotations.id AS id,
	CONCAT(employees.name, " ", employees.lastname) AS seller,
	customers.fullname AS customer,
	customers.doctype AS doctype,
	customers.docnum AS docnum,
	quotations.regdate AS issue,
	quotations.contact AS contact,
	CASE
		WHEN quotations.status = 'P' THEN 'Pendiente'
		WHEN quotations.status = 'A' THEN 'Aceptado'
		WHEN quotations.status = 'R' THEN 'Rechazado'
		ELSE 'Indefinido'
	END AS "status",
	quotations.currency AS currency,
	quotations.brand AS brand,
	quotations.model AS model,
	quotations.plate AS plate,
	quotations.serie AS serie,
	quotations.color AS color,
	quotations.total AS total,
	products.description AS product,
	quotedetails.description AS description,
	quotedetails.amount AS amount,
	products.unitprice
FROM quotations
JOIN customers ON quotations.customer_id = customers.id
JOIN employees ON quotations.employee_id = employees.id
JOIN quotedetails ON quotations.id = quotedetails.quotation_id
JOIN products ON quotedetails.product_id = products.id
ORDER BY quotedetails.id ASC;