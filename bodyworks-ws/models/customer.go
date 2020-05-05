package models

import (
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"log"
)
type Customer entities.Customer

func (c Customer) FindAll() ([]entities.Customer, error) {
	db := config.ConnectDB()
	defer db.Close()
	rs, err := db.Query("SELECT id, fullname, doctype, docnum, address FROM customers")
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rs.Close()
	var customers []entities.Customer
	for rs.Next()  {
		var customer entities.Customer
		rs.Scan(&customer.Id, &customer.Fullname, &customer.Doctype, &customer.Docnum, &customer.Adress)
		customers = append(customers, customer)
	}
	return customers, nil
}