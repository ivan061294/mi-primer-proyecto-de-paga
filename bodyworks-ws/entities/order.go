package entities

import (
	"time"
)
type OrderSubDetail struct {
	Id            int64 `json:"id"`
	Employee      int64 `json:"employee"`
	Workhours     int64  `json:"workhours"`
	Cloths        int64  `json:"cloths"`
}
type OrderDetail struct {
	Id             int64   `json:"id"`
	OrderId        int64   `json:"orderid"`
	Description    string  `json:"description"`
	Workhours      int64   `json:"workhours"`
	Cloths         int64   `json:"cloths"`
	SubDetail []OrderSubDetail `json:"subdetail"`
}
type Order struct {
	Id          int64 `json:"id"`
	QuoteId     int64 `json:"quoteid"`
	Worktype    string `json:"worktype"`
	Status      string `json:"status"`
	TotalHours  int64 `json:"totalhours"`
	TotalCloths int64 `json:"totalcloths"`
	StartDate   string `json:"startdate"`
	EndDate     string `json:"enddate"`
	Regdate     time.Time `json:"regdate"`
	Detail      []OrderDetail `json:"detail"`
	//Employees   []int64 `json:"employees"`
}
type OrderDetailView struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	WorkHours   int64  `json:"workhours"`
	Cloths      int64  `json:"cloths"`
	SubDetail []OrderSubDetail `json:"subdetail"`
}
type OrderView struct{
	Id       int64 `json:"id"`
	Seller   string `json:"seller"`
	Worktype string `json:"worktype"`
	Customer string `json:"customer"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Plate    string `json:"plate"`
	Serie    string `json:"serie"`
	Color    string `json:"color"`
	Doctype  string `json:"doctype"`
	Docnum      string            `json:"docnum"`
	Issue       time.Time         `json:"issue"`
	Contact     string            `json:"contact"`
	Status      string            `json:"status"`
	Currency    string            `json:"currency"`
	TotalHours  int64             `json:"totalhours"`
	TotalCloths int64             `json:"totalcloths"`
	StartDate   time.Time         `json:"startdate"`
	EndDate     time.Time         `json:"enddate"`
	Detail      []OrderDetailView `json:"detail"`
	//Employees   []int64           `json:"employees"`
}