CREATE VIEW viewinvoice AS
SELECT invoices.id AS id,
	CONCAT(employees.name, ' ' , employees.lastname) AS seller,
	customers.fullname AS customer,
	customers.doctype AS doctype,
	customers.docnum AS docnum,
	invoices.regdate AS issue,
	invoices.contact AS contact,
	CASE
		WHEN invoices.status = 'E' THEN 'Emitido'
		WHEN invoices.status = 'A' THEN 'Aprobado'
		WHEN invoices.status = 'O' THEN 'Observado'
		ELSE 'Indefinido'
	END AS status,
	invoices.currency AS currency,
	invoices.total AS total,
	invoices.observation AS observation,
	file1.name AS xmlsign,
	file2.name AS xmlsunat
FROM invoices
JOIN quotations ON invoices.quotation_id=quotations.id
JOIN customers ON invoices.customer_id=customers.id
JOIN employees ON invoices.employee_id=employees.id
LEFT JOIN documents file1 ON invoices.xmlfile_sign = file1.id
LEFT JOIN documents file2 ON invoices.xmlfile_sunat = file2.id
ORDER BY issue DESC;