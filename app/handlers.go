package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/rokafela/udemy-banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

/*
type Customer struct {
	Id      string `json:"customer_id" xml:"customerId"`
	Name    string `json:"full_name" xml:"fullName"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipCode"`
}

var Customers = []Customer{
	{Id: "1", Name: "Arief", City: "Jakarta", Zipcode: "12510"},
	{Id: "2", Name: "Darmawan", City: "Bogor", Zipcode: "16969"},
}
*/

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
