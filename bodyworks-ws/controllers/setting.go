package controllers

import (
	"net/http"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"gitlab.com/kevynestrada/bodywoks-ws/utils"
	"encoding/json"
)

func GetAllSetting(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var setting models.Setting
	settings, err := setting.FindAll()
	if err != nil {
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(settings)
}

func UpdateSetting(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	var setting models.Setting
	var response entities.Response
	response.Code = 0
	response.Message = "OK"
	if e := json.NewDecoder(r.Body).Decode(&setting); e != nil {
		http.Error(w, e.Error(), http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()
	_, err := setting.Update()
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
		http.Error(w, utils.ConvertObjToJson(response), http.StatusInternalServerError)
		return
	}
	http.Error(w, utils.ConvertObjToJson(response), http.StatusOK)
}

func AllowSetting(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, DELETE, PUT")
	if r.Method == "OPTIONS" {
		return
	}
}