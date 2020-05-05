package entities

type Product struct {
	Id int64 `json:"id"`
	Description string `json:"description"`
	Measurement string `json:"measurement"`
	Unitprice float64 `json:"unitprice"`
	Category string `json:"category"`
	Stock int64 `json:"stock"`
}