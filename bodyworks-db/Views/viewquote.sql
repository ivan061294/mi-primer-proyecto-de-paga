CREATE VIEW viewquote AS
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
	quotations.total AS total
FROM quotations
JOIN customers ON quotations.customer_id = customers.id
JOIN employees ON quotations.employee_id = employees.id
ORDER BY issue DESC;