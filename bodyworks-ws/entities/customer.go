package entities

type Customer struct {
	Id int64 `json:"id"`
	Fullname string `json:"fullname"`
	Doctype string `json:"doctype"`
	Docnum string `json:"docnum"`
	Adress string `json:"address"`
}