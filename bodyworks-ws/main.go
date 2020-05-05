package main

import (
	"github.com/gorilla/mux"
	"gitlab.com/kevynestrada/bodywoks-ws/controllers"
	"gitlab.com/kevynestrada/bodywoks-ws/utils"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/quotations", controllers.GetAllQuotation).Methods("GET")
	router.HandleFunc("/api/v1/quotations/{id}", controllers.GetQuotation).Methods("GET")
	router.HandleFunc("/api/v1/viewquotes/{id}", controllers.GetViewquote).Methods("GET")
	router.HandleFunc("/api/v1/quotations/{id}", controllers.UpdateQuotation).Methods("PUT")
	router.HandleFunc("/api/v1/quotations/{id}", controllers.DeleteQuotation).Methods("DELETE")
	router.HandleFunc("/api/v1/quotations", controllers.CreateQuotation).Methods("POST")
	router.HandleFunc("/api/v1/quotations/{id}", controllers.AllowQuotation).Methods("OPTIONS")

	router.HandleFunc("/api/v1/orders", controllers.GetAllOrder).Methods("GET")
	router.HandleFunc("/api/v1/orders/{id}", controllers.GetOrder).Methods("GET")
	router.HandleFunc("/api/v1/vieworders/{id}", controllers.GetVieworder).Methods("GET")
	router.HandleFunc("/api/v1/orders/{id}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/api/v1/orders/{id}", controllers.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/api/v1/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/v1/orders/{id}", controllers.AllowOrder).Methods("OPTIONS")

	router.HandleFunc("/api/v1/certifys", controllers.GetAllCertify).Methods("GET")
	router.HandleFunc("/api/v1/certifys/{id}", controllers.GetCertify).Methods("GET")
	router.HandleFunc("/api/v1/viewcertifys/{id}", controllers.GetViewCertify).Methods("GET")
	router.HandleFunc("/api/v1/certifys/{id}", controllers.UpdateCertify).Methods("PUT")
	router.HandleFunc("/api/v1/certifys/{id}", controllers.DeleteCertify).Methods("DELETE")
	router.HandleFunc("/api/v1/certifys", controllers.CreateCertify).Methods("POST")
	router.HandleFunc("/api/v1/certifys/{id}", controllers.AllowCertify).Methods("OPTIONS")

	router.HandleFunc("/api/v1/invoices", controllers.GetAllInvoice).Methods("GET")
	router.HandleFunc("/api/v1/invoices/{id}", controllers.GetInvoice).Methods("GET")
	router.HandleFunc("/api/v1/viewinvoices/{id}", controllers.GetViewinvoice).Methods("GET")
	router.HandleFunc("/api/v1/invoices/{id}", controllers.UpdateInvoice).Methods("PUT")
	router.HandleFunc("/api/v1/invoices/{id}", controllers.DeleteInvoice).Methods("DELETE")
	router.HandleFunc("/api/v1/invoices", controllers.CreateInvoice).Methods("POST")
	router.HandleFunc("/api/v1/invoices/{id}", controllers.AllowInvoice).Methods("OPTIONS")

	//SUNAT
	router.HandleFunc("/api/v2/invoices", controllers.CreateInvoiceSunat).Methods("GET")

	router.HandleFunc("/api/v1/customers", controllers.GetAllCustomer).Methods("GET")
	router.HandleFunc("/api/v1/products", controllers.GetAllProduct).Methods("GET")
	router.HandleFunc("/api/v1/employees", controllers.GetAllEmployee).Methods("GET")
	router.HandleFunc("/api/v1/supplies/quote/{id}", controllers.GetAllSupplieQuote).Methods("GET")

	router.HandleFunc("/api/v1/documents/{id}", controllers.GetFileContent).Methods("GET")

	router.HandleFunc("/api/v1/settings", controllers.GetAllSetting).Methods("GET")
	router.HandleFunc("/api/v1/settings", controllers.UpdateSetting).Methods("PUT")
	router.HandleFunc("/api/v1/settings", controllers.AllowSetting).Methods("OPTIONS")

	router.HandleFunc("/ws", controllers.ReportSale).Methods("GET")

	spa := utils.SpaHandler{StaticPath: "public", IndexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}