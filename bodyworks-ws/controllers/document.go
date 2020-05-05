package controllers

import (
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"strings"
	"net/http"
	"io"
	"log"
	"github.com/gorilla/mux"
)

func GetFileContent(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var doc models.Document
	doc.Name = mux.Vars(r)["id"]
	err := doc.GetContent()
	if err != nil {
		log.Print(err)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=" + doc.Name)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	io.Copy(w, strings.NewReader(string(doc.Content)))
}