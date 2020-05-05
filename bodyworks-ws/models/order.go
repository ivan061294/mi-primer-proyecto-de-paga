package models

import (
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"log"
	"time"
)

type OrderView entities.OrderView
type Order entities.Order
type OrderDetail entities.OrderDetail

func (q *Order) Count() (int64, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT COUNT(1) FROM orders"
	var countQuote int64
	err := db.QueryRow(query).Scan(&countQuote)
	if err != nil {
		return 0, err
	}
	return countQuote, nil
}
func (q *OrderView) FindAll() ([]OrderView, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT id, seller, customer, doctype, docnum, issue, contact, status, totalhours, totalcloths, startdate, enddate FROM vieworder"
	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var orders []OrderView
	orders = []OrderView{}
	for rs.Next() {
		err := rs.Scan(&q.Id, &q.Seller, &q.Customer, &q.Doctype, &q.Docnum, &q.Issue, &q.Contact, &q.Status, &q.TotalHours, &q.TotalCloths, &q.StartDate, &q.EndDate)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *q)
	}
	return orders, nil
}
func (q *Order) FindOne(id int64) (Order, error) {
	db := config.ConnectDB()
	defer db.Close()
	var queryHead = "SELECT id, quotation_id, worktype, status, totalhours, totalcloths, startdate, enddate, regdate FROM orders WHERE id=?"
	var queryBody = "SELECT id, order_id, description, workhours, cloths FROM orderdetails WHERE order_id=?"
	var querySubdetail = "SELECT id, employee_id, workhours, cloths FROM orderdetailemployees WHERE orderdetail_id=?"
	db.QueryRow(queryHead, id).Scan(&q.Id, &q.QuoteId, &q.Worktype, &q.Status, &q.TotalHours, &q.TotalCloths, &q.StartDate, &q.EndDate, &q.Regdate)
	rs0, err := db.Query(queryBody, id)
	if err != nil {
		return *q, err
	}
	for rs0.Next() {
		var detail entities.OrderDetail
		err := rs0.Scan(&detail.Id, &detail.OrderId, &detail.Description, &detail.Workhours, &detail.Cloths)
		if err != nil {
			return *q, err
		}
		rs1, err := db.Query(querySubdetail, detail.Id)
		if err != nil {
			return *q, err
		}
		for rs1.Next() {
			var subdetail entities.OrderSubDetail
			err := rs1.Scan(&subdetail.Id, &subdetail.Employee, &subdetail.Workhours, &subdetail.Cloths)
			if err != nil {
				return *q, err
			}
			detail.SubDetail = append(detail.SubDetail, subdetail)
		}
		q.Detail = append(q.Detail, detail)
	}
	return *q, nil
}
func (q *OrderView) FindOne(id int64) (OrderView, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT worktype, customer, issue, brand, model, plate, serie, color, totalhours, totalcloths, startdate, enddate, detailid, description, workhours, cloths FROM vieworderdetail WHERE id=?"
	var subdatail = "SELECT employee_id, workhours, cloths FROM orderdetailemployees WHERE orderdetail_id=?"
	rs, err := db.Query(query, id)
	if err != nil {
		return *q, err
	}
	defer rs.Close()
	for rs.Next() {
		var detail entities.OrderDetailView
		err := rs.Scan(&q.Worktype ,&q.Customer, &q.Issue, &q.Brand, &q.Model, &q.Plate, &q.Serie, &q.Color, &q.TotalHours, &q.TotalCloths, &q.StartDate, &q.EndDate, &detail.Id, &detail.Description, &detail.WorkHours, &detail.Cloths)
		if err != nil {
			return *q, err
		}
		rs1, err := db.Query(subdatail, detail.Id)
		if err != nil {
			return *q, err
		}
		for rs1.Next() {
			var subdetail entities.OrderSubDetail
			err := rs1.Scan(&subdetail.Employee, &subdetail.Workhours, &subdetail.Cloths)
			if err != nil {
				return *q, err
			}
			detail.SubDetail = append(detail.SubDetail, subdetail)
		}
		defer rs1.Close()
		q.Detail = append(q.Detail, detail)
	}
	return *q, nil
}
func (q *Order) Create() (interface{}, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "INSERT INTO orders (id, quotation_id, worktype, status, totalhours, totalcloths, startdate, enddate, regdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	startdate, _ := time.Parse("02/01/2006", q.StartDate)
	enddate, _ := time.Parse("02/01/2006", q.EndDate)
	tx, err := db.Begin()
	rs, err := tx.Exec(query, q.Id, q.QuoteId, q.Worktype, "E", q.TotalHours, q.TotalCloths, startdate, enddate, time.Now())
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	q.Id, err = rs.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, detail := range q.Detail {
		rs, err := tx.Exec("INSERT INTO orderdetails (order_id, description, workhours, cloths) VALUES (?, ?, ?, ?)", q.Id, detail.Description, detail.Workhours, detail.Cloths)
		orderDetailId, err := rs.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		for _, subdetail := range detail.SubDetail {
			tx.Exec("INSERT INTO orderdetailemployees (orderdetail_id, employee_id, workhours, cloths) VALUES (?, ?, ?, ?)", orderDetailId, subdetail.Employee, subdetail.Workhours, subdetail.Cloths)
		}
	}
	tx.Commit()
	return q, nil
}
func (q *Order) Delete(id int64) (entities.Response) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "call sp_delete_order(?, @coderror, @msgerror);"
	var response entities.Response
	err := db.QueryRow(query, id).Scan(&response.Code, &response.Message)
	if err != nil {
		log.Print(err)
	}
	return response
}