package models

import (
	"database/sql"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
)

type QuoteView entities.QuoteView
type QuoteDetailView entities.QuoteDetailView
type Quotation entities.Quotation
type QuoteDetail entities.QuoteDetail

func (q *Quotation) Count() (int64, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT COUNT(1) FROM quotations"
	var countQuote int64
	err := db.QueryRow(query).Scan(&countQuote)
	if err != nil {
		return 0, err
	}
	return countQuote, nil
}
func (q *QuoteView) FindAll() ([]QuoteView, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT id, seller, customer, doctype, docnum, issue, contact, status, currency, total FROM viewquote"
	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var quoteviews []QuoteView
	quoteviews = []QuoteView{}
	for rs.Next() {
		err := rs.Scan(&q.Id, &q.Seller, &q.Customer, &q.Doctype, &q.Docnum, &q.Issue, &q.Contact, &q.Status, &q.Currency, &q.Total)
		if err != nil {
			return nil, err
		}
		quoteviews = append(quoteviews, *q)
	}
	return quoteviews, nil
}
func (q *QuoteView) FindOne(id int64) (QuoteView, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT customer, contact, seller, currency, brand, model, plate, serie, color, product, description, amount, unitprice FROM viewquotedetail WHERE id=?"
	rs, err := db.Query(query, id)
	if err != nil {
		return *q, err
	}
	defer rs.Close()
	for rs.Next() {
		var brand sql.NullString
		var model sql.NullString
		var plate sql.NullString
		var serie sql.NullString
		var color sql.NullString
		var detail entities.QuoteDetailView
		err := rs.Scan(&q.Customer, &q.Contact, &q.Seller, &q.Currency, &brand, &model, &plate, &serie, &color, &detail.Product, &detail.Description, &detail.Amount, &detail.Unitprice)
		if err != nil {
			return *q, err
		}
		if brand.Valid {
			q.Brand = brand.String
		}
		if model.Valid {
			q.Model = model.String
		}
		if plate.Valid {
			q.Plate = plate.String
		}
		if serie.Valid {
			q.Serie = serie.String
		}
		if color.Valid {
			q.Color = color.String
		}
		q.Detail = append(q.Detail, detail)
	}
	return *q, nil
}
func (q *Quotation) FindOne(id int64) (Quotation, error) {
	db := config.ConnectDB()
	defer db.Close()
	var queryHead = "SELECT id, employee_id, customer_id, contact, status, currency, total, regdate, brand, model, plate, serie, color FROM quotations WHERE id=?"
	var queryBody = "SELECT id, quotation_id, product_id, description, amount, price FROM quotedetails WHERE quotation_id=?"
	db.QueryRow(queryHead, id).Scan(&q.Id, &q.EmployeeId, &q.CustomerId, &q.Contact, &q.Status, &q.Currency, &q.Total, &q.Regdate, &q.Brand, &q.Model, &q.Plate, &q.Serie, &q.Color)
	rs, err := db.Query(queryBody, id)
	if err != nil {
		log.Print(err)
		return *q, err
	}
	defer rs.Close()
	for rs.Next() {
		var detail entities.QuoteDetail
		err := rs.Scan(&detail.Id, &detail.QuotationId, &detail.ProductId, &detail.Description, &detail.Amount, &detail.Price)
		if err != nil {
			log.Print(err)
			return *q, err
		}
		q.Detail = append(q.Detail, detail)
	}
	return *q, nil
}

func (q *Quotation) Create() (interface{}, error) {
	var query = "INSERT INTO quotations (employee_id, customer_id, contact, status, currency, total, regdate, brand, model, plate, serie, color) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	var queryDetail = "INSERT INTO quotedetails (quotation_id, product_id, description, amount, price) VALUES (?, ?, ?, ?, ?)"

	db := config.ConnectDB()
	defer db.Close()
	tx, err := db.Begin()
	rs, err := tx.Exec(query, q.EmployeeId, q.CustomerId, q.Contact, "P", q.Currency, q.Total, time.Now(), q.Brand, q.Model, q.Plate, q.Serie, q.Color)
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

func (q *QuoteDetail) Create() (interface{}, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "INSERT INTO quotedetails (quotation_id, product_id, description, amount, price) VALUES (?, ?, ?, ?, ?)"
	rs, err := db.Exec(query, q.QuotationId, q.ProductId, q.Description, q.Amount, q.Price)
	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return nil, err
	}
	return id, nil
}
func (q *QuoteDetail) Delete(id int64) error {
	db := config.ConnectDB()
	defer db.Close()
	var query = "DELETE FROM quotedetails WHERE quotation_id=?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (q *Quotation) Delete(id int64) error {
	db := config.ConnectDB()
	defer db.Close()
	var query = "DELETE FROM quotations WHERE id=?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (q *Quotation) IsPending(id int64) bool {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT status FROM quotations WHERE id=?"
	var status string
	db.QueryRow(query, id).Scan(&status)
	if status == "P" {
		return true
	} else {
		return false
	}
}
func (q *Quotation) Update(id int64) (Quotation, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "UPDATE quotations SET employee_id = ?, customer_id = ?, contact = ?, status = ?, currency = ?, total = ?, brand = ?, model = ?, plate = ?, serie = ?, regdate = ? WHERE id = ?"
	var update = "UPDATE quotedetails SET product_id = ?, description = ?, amount = ?, price = ? WHERE id = ?"
	var create = "INSERT INTO quotedetails (quotation_id, product_id, description, amount, price) VALUES (?, ?, ?, ?, ?)"
	tx, err := db.Begin()
	_, err = tx.Exec(query, q.EmployeeId, q.CustomerId, q.Contact, q.Status, q.Currency, q.Total, q.Brand, q.Model, q.Plate, q.Serie, q.Regdate, id)
	if err != nil {
		tx.Rollback()
		return *q, err
	}
	var aID = make([]int64, 0)
	aID = append(aID, q.Id)
	for _, detail := range q.Detail {
		var err error
		if detail.Id > 0 {
			_, err = tx.Exec(update, detail.ProductId, detail.Description, detail.Amount, detail.Price, detail.Id)
		} else {
			var rs sql.Result
			rs, err = tx.Exec(create, id, detail.ProductId, detail.Description, detail.Amount, detail.Price)
			detail.Id, _ = rs.LastInsertId()
		}
		if err != nil {
			tx.Rollback()
			return *q, err
		}
		aID = append(aID, detail.Id)
	}
	args := make([]interface{}, len(aID))
	for i, id := range aID {
		args[i] = id
	}
	var clean = "DELETE FROM quotedetails WHERE quotation_id = ? AND id NOT IN (?" + strings.Repeat(",?", len(args)-2) + ")"
	_, err = tx.Exec(clean, args...)
	if err != nil {
		tx.Rollback()
		return *q, err
	}
	tx.Commit()
	return *q, nil
}

/*
func (q *QuoteDetail) Update(id int64) (QuoteDetail, error) {
	db := config.ConnectDB()
	defer db.Close()
	var update = "UPDATE quotedetails SET product_id = ?, description = ?, amount = ?, price = ? WHERE id = ?"
	var create = "INSERT INTO quotedetails (quotation_id, product_id, description, amount, price) VALUES (?, ?, ?, ?, ?)"
	var err error
	if q.Id != 0 {
		_, err = db.Exec(update, q.ProductId, q.Description, q.Amount, q.Price, q.Id)
	} else {
		_, err = db.Exec(create, id, q.ProductId, q.Description, q.Amount, q.Price)
	}
	if err != nil {
		return *q, err
	}
	return *q, nil
}
*/
