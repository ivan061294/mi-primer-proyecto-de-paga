package controllers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type Message struct {
	CountQuote int64 `json:"countQuote"`
	CountOrder int64 `json:"countOrder"`
	CountCertify int64 `json:"countCertify"`
	CountInvoice int64 `json:"countInvoice"`
	SaleForMonth []entities.Salemonth `json:"saleForMonth"`
}

var clients = make(map[*websocket.Conn]bool)

func ReportSale(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrade(w, r)
	if err != nil {
		log.Print(err)
	}
	clients[ws] = true
	go Writer(ws)
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return ws, err
	}
	return ws, nil
}
func Writer(ws *websocket.Conn)  {
	var reportQuote models.Quotation
	var reportOrder models.Order
	var reportInvoice models.Invoice

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, ws)
			return
		}
		log.Print(string(p))
		countQuote, err := reportQuote.Count()
		if err != nil {
			log.Print(err)
		}
		countOrder, err := reportOrder.Count()
		if err != nil {
			log.Print(err)
		}
		countInvoice, err := reportInvoice.Count()
		if err != nil {
			log.Print(err)
		}
		salemonths, err := reportInvoice.GetSaleForMonth()
		if err != nil {
			log.Print(err)
		}

		var reportSale Message
		reportSale.CountQuote = countQuote
		reportSale.CountOrder = countOrder
		reportSale.CountInvoice = countInvoice
		reportSale.SaleForMonth = salemonths
		message, err := json.Marshal(reportSale)
		if err != nil {
			log.Print(err)
		}

		for ws := range clients {
			if err := ws.WriteMessage(messageType, message); err != nil {
				log.Println(err)
				delete(clients, ws)
				return
			}
		}
	}
}