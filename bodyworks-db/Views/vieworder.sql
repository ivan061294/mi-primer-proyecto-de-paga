CREATE VIEW vieworder AS
SELECT orders.id AS id,
	CONCAT(employees.name, ' ' , employees.lastname) AS seller,
	customers.fullname AS customer,
	customers.doctype AS doctype,
	customers.docnum AS docnum,
	orders.regdate AS issue,
	quotations.contact AS contact,
	CASE
		WHEN orders.status = 'E' THEN 'En proceso'
		WHEN orders.status = 'T' THEN 'Terminado'
		ELSE 'Indefinido'
	END AS status,
	orders.totalhours AS totalhours,
	orders.totalcloths AS totalcloths,
	orders.startdate AS startdate,
	orders.enddate AS enddate
FROM orders
JOIN quotations ON orders.quotation_id=quotations.id
JOIN customers ON quotations.customer_id=customers.id
JOIN employees ON quotations.employee_id=employees.id
ORDER BY issue DESC;