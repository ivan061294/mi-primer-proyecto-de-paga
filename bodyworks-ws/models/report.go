package models

import (
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
)

func (q *Invoice) GetSaleForMonth() ([]entities.Salemonth, error) {
	var salemonths []entities.Salemonth
	var salemonth entities.Salemonth
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT month, ventas FROM viewsaleformonth"
	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rs.Next() {
		err := rs.Scan(&salemonth.Month, &salemonth.Sale)
		if err != nil {
			return nil, err
		}
		salemonths = append(salemonths, salemonth)
	}
	return salemonths, nil
}