package entities

import "time"

type CertyView struct {
	Id          int64     `json:"id"`
	Customer    string    `json:"customer"`
	Contact     string    `json:"contact"`
	Area        string    `json:"area"`
	Brand       string    `json:"brand"`
	Model       string    `json:"model"`
	Plate       string    `json:"plate"`
	Employee    string    `json:"employe"`
	QuoteId     int64     `json:"quoteId"`
	OrderId     int64     `json:"orderId"`
	IniDate     time.Time `json:"inidate"`
	FinDate     time.Time `json:"findate"`
	Description string    `json:"description"`
	Observation string    `json:"observation"`
}

type Certify struct {
	Id           int64     `json:"id"`
	QuotationId  int64     `json:"quotationId"`
	OrderId      int64     `json:"orderId"`
	Regdate      time.Time `json:"regdate"`
	Description  string    `json:"description"`
}