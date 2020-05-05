package models

import (
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"log"
	"time"
)


type Invoice entities.Invoice
type InvoiceView entities.InvoiceView

func (q *Invoice) Count() (int64, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT COUNT(1) FROM invoices"
	var countQuote int64
	err := db.QueryRow(query).Scan(&countQuote)
	if err != nil {
		return 0, err
	}
	return countQuote, nil
}
func (q *InvoiceView) FindAll() ([]InvoiceView, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = `SELECT id,
						seller,
						customer,
						doctype,
						docnum,
						issue,
						contact,
						status,
						currency,
						total,
						IFNULL(observation,''),
						IFNULL(xmlsign,''),
						IFNULL(xmlsunat,'')
					FROM viewinvoice`
	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var invoices []InvoiceView
	invoices = []InvoiceView{}
	for rs.Next() {
		err := rs.Scan(&q.Id,
			&q.Seller, &q.Customer, &q.Doctype, &q.Docnum,
			&q.Issue, &q.Contact, &q.Status, &q.Currency,
			&q.Total, &q.Observation, &q.Xmlsign, &q.Xmlsunat)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, *q)
	}
	return invoices, nil
}

func (q *Invoice) Create() (interface{}, error) {
	var query = "INSERT INTO invoices (employee_id, customer_id, quotation_id, contact, status, currency, total, regdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	var queryDetail = "INSERT INTO quotedetails (quotation_id, product_id, description, amount, price) VALUES (?, ?, ?, ?, ?)"

	db := config.ConnectDB()
	defer db.Close()
	tx, err := db.Begin()
	rs, err := tx.Exec(query, q.EmployeeId, q.CustomerId, q.QuotationId, q.Contact, "E", q.Currency, q.Total, time.Now())
	if err != nil {
		return nil, err
	}
	q.Id, err = rs.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for idx := range q.Detail {
		rs, err := tx.Exec(queryDetail, q.Id, q.Detail[idx].ProductId, q.Detail[idx].Description, q.Detail[idx].Amount, q.Detail[idx].Price)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		id, err := rs.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		q.Detail[idx].Id = id
	}
	tx.Commit()
	return q, nil
}
func (q *Invoice) Delete(id int64) (entities.Response) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "call sp_delete_invoice(?, @coderror, @msgerror);"
	var response entities.Response
	err := db.QueryRow(query, id).Scan(&response.Code, &response.Message)
	if err != nil {
		log.Print(err)
	}
	return response
}