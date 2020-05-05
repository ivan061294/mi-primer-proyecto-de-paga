package models

import (
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
)

type Document entities.Document

func (d *Document) GetContent() error {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT content FROM documents WHERE name=?"
	var content []byte
	err := db.QueryRow(query, d.Name).Scan(&content)
	if err != nil {
		return err
	}
	d.Content = content
	return nil
}