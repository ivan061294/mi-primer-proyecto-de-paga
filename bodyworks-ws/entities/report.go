package entities

import "time"

type Salemonth struct {
	Month time.Time `json:"month"`
	Sale  float64   `json:"sale"`
}