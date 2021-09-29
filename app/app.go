package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	"github.com/rokafela/udemy-banking/domain"
	"github.com/rokafela/udemy-banking/logger"
	"github.com/rokafela/udemy-banking/service"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mandatory_env := []string{
		"APP_ADDRESS",
		"APP_PORT",
		"DB_ADDRESS",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
	}

	for _, v := range mandatory_env {
		_, db_address_bool := os.LookupEnv(v)

		if !db_address_bool {
			logger.Fatal("Environment variable not defined|" + v)
		}
	}
}

func Start() {
	// router
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewRandomStub())}

	// routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomerById).Methods(http.MethodGet)

	// server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("APP_ADDRESS"), os.Getenv("APP_PORT")), router))
}
