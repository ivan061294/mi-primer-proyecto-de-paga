package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"gitlab.com/kevynestrada/bodywoks-ws/utils"
	"log"
	"net/http"
	"strconv"
)

func GetAllOrder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var orderView models.OrderView
	orders, err := orderView.FindAll()
	if len(orders) == 0 {
		log.Print("No se encontro ninguna orden de trabajo")
	}
	if err != nil {
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(orders)
}
func GetOrder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var order models.Order
	orders, err := order.FindOne(id)
	if err != nil {
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(orders)
}
func GetVieworder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var order models.OrderView
	orderview, err := order.FindOne(id)
	if err != nil {
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	log.Print(orderview)
	json.NewEncoder(w).Encode(orderview)
}
func CreateOrder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var order models.Order
	var response entities.Response
	response.Code = 0
	response.Message = "OK"
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		response.Code = 1
		response.Message = err.Error()
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(response), http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()
	log.Print(utils.ConvertObjToJson(order))
	id, err := order.Create()
	if err != nil {
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(id)
}
func UpdateOrder(w http.ResponseWriter, r *http.Request)  {

}
func DeleteOrder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var order models.Order
	response := order.Delete(id)
	http.Error(w, utils.ConvertObjToJson(response), http.StatusOK)
}
func AllowOrder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, DELETE, PUT")
	log.Print("AllowQuotation")
	if r.Method == "OPTIONS" {
		return
	}
}