package models

import (
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"time"
)

type CertyView entities.CertyView
type Certify entities.Certify

func (c *CertyView) FindAll() ([]CertyView, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT id, customer, contact, brand, model, plate, seller, description, observation FROM viewcerty"
	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	certiviews := []CertyView{}
	for rs.Next() {
		err := rs.Scan(&c.Id, &c.Customer, &c.Contact, &c.Brand, &c.Model, &c.Plate, &c.Employee, &c.Description, &c.Observation)
		if err != nil {
			return nil, err
		}
		certiviews = append(certiviews, *c)
	}
	return certiviews, nil
}

func (c Certify) Create() (Certify, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "INSERT INTO certifys (quotation_id, order_id, regdate, description) VALUES (?, ?, ?, ?)"
	rs, err := db.Exec(query, c.QuotationId, c.OrderId, time.Now(), c.Description)
	if err != nil {
		return c, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return c, err
	}
	c.Id = id
	return c, nil
}