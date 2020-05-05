package entities

import "time"

type QuoteDetail struct {
	Id          int64   `json:"id"`
	QuotationId int64   `json:"quotationid"`
	ProductId   int64   `json:"product"`
	Description string  `json:"description"`
	Amount      int64   `json:"quantity"`
	Price       float64 `json:"unitprice"`
}
type Quotation struct {
	Id         int64         `json:"id"`
	EmployeeId int64         `json:"employee"`
	CustomerId int64         `json:"customer"`
	Contact    string        `json:"contact"`
	Status     string        `json:"status"`
	Currency   string        `json:"currency"`
	Total      float64       `json:"total"`
	Brand      string        `json:"brand"`
	Model      string        `json:"model"`
	Plate      string        `json:"plate"`
	Serie      string        `json:"serie"`
	Color      string        `json:"color"`
	Regdate    time.Time     `json:"regdate"`
	Detail     []QuoteDetail `json:"detail"`
}
type QuoteDetailView struct {
	//Id          int64  `json:"id"`
	//Quoteid     int64  `json:"quoteid"`
	Product     string  `json:"product"`
	Description string  `json:"description"`
	Amount      int64   `json:"amount"`
	Unitprice   float64 `json:"unitprice"`
}
type QuoteView struct {
	Id       int64             `json:"id"`
	Seller   string            `json:"seller"`
	Customer string            `json:"customer"`
	Doctype  string            `json:"doctype"`
	Docnum   string            `json:"docnum"`
	Issue    time.Time         `json:"issue"`
	Contact  string            `json:"contact"`
	Status   string            `json:"status"`
	Currency string            `json:"currency"`
	Brand    string            `json:"brand"`
	Model    string            `json:"model"`
	Plate    string            `json:"plate"`
	Serie    string            `json:"serie"`
	Color    string            `json:"color"`
	Total    float64           `json:"total"`
	Detail   []QuoteDetailView `json:"detail"`
}
