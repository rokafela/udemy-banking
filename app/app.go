package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rokafela/udemy-banking/domain"
	"github.com/rokafela/udemy-banking/service"
)

func Start() {
	// router
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewRandomStub())}

	// routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	// .Queries("status", "{status:(?:active|inactive)}")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomerById).Methods(http.MethodGet)

	// server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
