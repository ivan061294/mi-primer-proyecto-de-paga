CREATE VIEW viewcerty AS
SELECT
	certifys.id AS id,
	customers.fullname AS customer,
    quotations.contact AS contact,
    IFNULL(quotations.brand, '') AS brand,
    IFNULL(quotations.model, '') AS model,
    IFNULL(quotations.plate, '') AS plate,
    CONCAT(employees.name, ' ' , employees.lastname) AS seller,
    IFNULL(certifys.description, '') AS description,
    IFNULL(certifys.observation, '') AS observation
FROM certifys
JOIN quotations ON certifys.quotation_id = quotations.id
JOIN customers ON quotations.customer_id = customers.id
JOIN employees ON quotations.employee_id = employees.id
ORDER BY certifys.regdate DESC;