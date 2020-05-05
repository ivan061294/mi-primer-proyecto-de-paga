package controllers

import (
	"encoding/json"
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"gitlab.com/kevynestrada/bodywoks-ws/utils"
	"log"
	"net/http"
)

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var employee models.Employee
	employees, err := employee.FindAll()
	if err != nil {
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	if len(employees) == 0 {
		log.Print("No se encontro ningun empleado")
		http.Error(w, "", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(employees)
}