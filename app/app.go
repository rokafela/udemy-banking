package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
	"github.com/rokafela/udemy-banking/domain"
	"github.com/rokafela/udemy-banking/logger"
	"github.com/rokafela/udemy-banking/service"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Panic("Error loading .env file")
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
			logger.Panic("Environment variable not defined|" + v)
		}
	}
}

func Start() {
	// router
	router := mux.NewRouter()

	// repository initialization
	dbPool := createDbPool()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbPool)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbPool)
	transactionRepositoryDb := domain.SetTransactionRepositoryDb(dbPool)

	// handler initialization
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandlers{service.NewAccountService(accountRepositoryDb)}
	th := TransactionHandler{service.SetTransactionService(transactionRepositoryDb)}

	// routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.CreateNewAccount).Methods(http.MethodPost)

	router.HandleFunc("/transaction", th.CreateTransaction).Methods(http.MethodPost)

	// server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("APP_ADDRESS"), os.Getenv("APP_PORT")), router))
}

func createDbPool() *sqlx.DB {
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_address := os.Getenv("DB_ADDRESS")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dbProperties := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_address, db_port, db_name)

	dbPool, err := sqlx.Connect("mysql", dbProperties)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// See "Important settings" section.
	dbPool.SetConnMaxLifetime(time.Minute * 3)
	dbPool.SetMaxOpenConns(10)
	dbPool.SetMaxIdleConns(10)

	return dbPool
}
