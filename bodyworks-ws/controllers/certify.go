package controllers

import (
	"net/http"
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"gitlab.com/kevynestrada/bodywoks-ws/utils"
	"log"
	"encoding/json"
)

func GetAllCertify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var certyView models.CertyView
	certys, err := certyView.FindAll()
	if len(certys) == 0 {
		log.Print("No se encontro ningun acta de entrega")
	}
	if err != nil {
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(certys)
}
func GetCertify(w http.ResponseWriter, r *http.Request) {}
func GetViewCertify(w http.ResponseWriter, r *http.Request) {}
func UpdateCertify(w http.ResponseWriter, r *http.Request) {}
func DeleteCertify(w http.ResponseWriter, r *http.Request) {}
func CreateCertify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var certify models.Certify
	if e := json.NewDecoder(r.Body).Decode(&certify); e != nil {
		http.Error(w, e.Error(), http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()
	id, err := certify.Create()
	if err != nil {
		http.Error(w, utils.ConvertObjToJson(err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(id)
}
func AllowCertify(w http.ResponseWriter, r *http.Request) {}