package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rokafela/udemy-banking/app/modules/customer"
	"github.com/rokafela/udemy-banking/app/modules/greet"
	"github.com/rokafela/udemy-banking/app/modules/time"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/greet", greet.Greet).Methods(http.MethodGet)

	router.HandleFunc("/customers", customer.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", customer.CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]*}", customer.GetCustomerById).Methods(http.MethodGet)

	router.HandleFunc("/api/time", time.GetCurrentTime).Methods(http.MethodGet)

	// server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
