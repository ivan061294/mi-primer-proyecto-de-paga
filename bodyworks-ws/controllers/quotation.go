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

func GetAllQuotation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var quoteView models.QuoteView
	quotations, err := quoteView.FindAll()
	if err != nil {
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	if len(quotations) == 0 {
		log.Print("No se encontro ninguna cotizacion")
		//http.Error(w, "[]", http.StatusNotFound)
		//	return
	}
	json.NewEncoder(w).Encode(quotations)
}
func GetQuotation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var quote models.Quotation
	quotation, err := quote.FindOne(id)
	if err != nil {
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(quotation)
}
func GetViewquote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var quote models.QuoteView
	quotation, err := quote.FindOne(id)
	if err != nil {
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(quotation)
}
func CreateQuotation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var quotation models.Quotation
	if e := json.NewDecoder(r.Body).Decode(&quotation); e != nil {
		http.Error(w, e.Error(), http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()
	id, err := quotation.Create()
	if err != nil {
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(id)
}
func UpdateQuotation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var quotation models.Quotation
	var response entities.Response
	response.Code = 0
	response.Message = "OK"
	if e := json.NewDecoder(r.Body).Decode(&quotation); e != nil {
		http.Error(w, e.Error(), http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()
	_, err := quotation.Update(id)
	if err != nil {
		log.Print(err)
		response.Code = 1
		response.Message = err.Error()
		http.Error(w, utils.ConvertObjToJson(response), http.StatusInternalServerError)
		return
	}
	http.Error(w, utils.ConvertObjToJson(response), http.StatusOK)
}
func DeleteQuotation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	log.Print("DeleteQuotation")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var quotation models.Quotation
	var quoteDetail models.QuoteDetail
	var response entities.Response
	response.Code = 0
	response.Message = "OK"
	isPending := quotation.IsPending(id)
	if !isPending {
		response.Code = 1
		response.Message = utils.MsgErrorNoValidStatus
		log.Print(utils.ConvertObjToJson(response))
		http.Error(w, utils.ConvertObjToJson(response), http.StatusOK)
		return
	}
	err := quotation.Delete(id)
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
		log.Print(utils.ConvertObjToJson(response))
		http.Error(w, utils.ConvertObjToJson(response), http.StatusInternalServerError)
		return
	}
	err2 := quoteDetail.Delete(id)
	if err2 != nil {
		response.Code = 1
		response.Message = err2.Error()
		log.Print(utils.ConvertObjToJson(response))
		http.Error(w, utils.ConvertObjToJson(response), http.StatusInternalServerError)
		return
	}
	log.Print(utils.ConvertObjToJson(response))
	http.Error(w, utils.ConvertObjToJson(response), http.StatusOK)
}
func AllowQuotation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, DELETE, PUT")
	if r.Method == "OPTIONS" {
		return
	}
}
