package models

import (
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"log"
)

type Product entities.Product

func (product Product) FindAll() ([]entities.Product, error)  {
	db := config.ConnectDB()
	defer db.Close()
	rs, err := db.Query("SELECT id, description, measurement, unitprice, category, stock FROM products")
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rs.Close()
	var products []entities.Product
	for rs.Next() {
		var product entities.Product
		rs.Scan(&product.Id, &product.Description, &product.Measurement, &product.Unitprice, &product.Category, &product.Stock)
		products = append(products, product)
	}
	return products, nil
}