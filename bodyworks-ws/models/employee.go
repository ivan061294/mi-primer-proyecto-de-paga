package models

import (
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
)

type Employee entities.Employee

func (o *Employee) FindAll() ([]Employee, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT id, name, lastname, care FROM employees"
	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var employees []Employee
	for rs.Next() {
		err := rs.Scan(&o.Id, &o.Name, &o.Lastname, &o.Care)
		if err != nil {
			return nil, err
		}
		employees = append(employees, *o)
	}
	return employees, nil
}