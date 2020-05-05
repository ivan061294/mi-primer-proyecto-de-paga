package controllers

import (
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"gitlab.com/kevynestrada/bodywoks-ws/utils"
	"encoding/json"
	"net/http"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var product models.Product
	products, err := product.FindAll()
	if len(products) == 0 {
		http.Error(w, utils.ConvertObjToJson(products), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, utils.ConvertObjToJson(products), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}