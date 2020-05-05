package models

import (
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
)

type Supplie entities.Supplie

func (s *Supplie) FindPerQuote(id int64) ([]Supplie, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT quotation_id, supplie_id, autoamount FROM quotesupplieview WHERE quotation_id=?"
	rs, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var supplies []Supplie
	for rs.Next() {
		err := rs.Scan(&s.Quote, &s.Product, &s.Amount)
		if err != nil {
			return nil, err
		}
		supplies = append(supplies, *s)
	}
	return supplies, nil
}