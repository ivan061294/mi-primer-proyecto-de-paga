package controllers

import (
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"encoding/json"
	"net/http"
)

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var customer models.Customer
	customers, err := customer.FindAll()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(customers)
	}
}