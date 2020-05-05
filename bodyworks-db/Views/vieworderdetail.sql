CREATE VIEW vieworderdetail AS
SELECT orders.id AS id,
	orders.worktype AS worktype,
	CONCAT(employees.name, ' ' , employees.lastname) AS seller,
	customers.fullname AS customer,
	customers.doctype AS doctype,
	customers.docnum AS docnum,
	orders.regdate AS issue,
	IFNULL(quotations.contact, '') AS contact,
	IFNULL(quotations.brand, '') AS brand,
	IFNULL(quotations.model, '') AS model,
	IFNULL(quotations.plate, '') AS plate,
	IFNULL(quotations.serie, '') AS serie,
	IFNULL(quotations.color, '') AS color,
	CASE
		WHEN orders.status = 'E' THEN 'En proceso'
		WHEN orders.status = 'T' THEN 'Terminado'
		ELSE 'Indefinido'
	END AS status,
	orders.totalhours AS totalhours,
	orders.totalcloths AS totalcloths,
	orders.startdate AS startdate,
	orders.enddate AS enddate,
	orderdetails.id AS detailid,
	orderdetails.description AS description,
	orderdetails.workhours AS workhours,
	orderdetails.cloths AS cloths
FROM orders
JOIN quotations ON orders.quotation_id=quotations.id
JOIN customers ON quotations.customer_id=customers.id
JOIN employees ON quotations.employee_id=employees.id
JOIN orderdetails ON orders.id=orderdetails.order_id;