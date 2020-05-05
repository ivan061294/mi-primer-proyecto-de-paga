package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"gitlab.com/kevynestrada/bodywoks-ws/utils"
	"log"
	"net/http"
	"strconv"
)

func GetAllSupplieQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var supplie models.Supplie
	supplies, err := supplie.FindPerQuote(id)
	if err != nil {
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	if len(supplies) == 0 {
		log.Print("No se encontro ningun insumo")
		http.Error(w, "", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(supplies)
}