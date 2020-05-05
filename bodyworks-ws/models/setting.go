package models

import (
	"gitlab.com/kevynestrada/bodywoks-ws/config"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"time"
)

type Setting entities.Setting

func (q *Setting) FindAll() ([]Setting, error) {
	db := config.ConnectDB()
	defer db.Close()
	var query = "SELECT name, value FROM settings"
	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	settings := []Setting{}
	for rs.Next() {
		err := rs.Scan(&q.Name, &q.Value)
		if err != nil {
			return nil, err
		}
		settings = append(settings, *q)
	}
	return settings, nil
}

func (s *Setting) Update() (Setting, error) {
	db := config.ConnectDB()
	defer db.Close()
	var update = "UPDATE settings SET value = ?, regdate = ? WHERE name = ?"
	_, err := db.Exec(update, s.Value, time.Now(), s.Name)
	if err != nil {
		return *s, err
	}
	return *s, nil
}