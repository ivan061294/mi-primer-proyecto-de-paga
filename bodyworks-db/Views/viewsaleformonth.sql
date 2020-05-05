CREATE VIEW viewsaleformonth AS
SELECT
	DATE_FORMAT(regdate, '%Y-%m-01') AS month,
	SUM(total) AS ventas
FROM invoices
WHERE regdate >= last_day(now()) + interval 1 day - interval 7 month
GROUP BY DATE_FORMAT(regdate, '%Y-%m-01');