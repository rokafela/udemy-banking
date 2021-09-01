package customer

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"fullName"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipCode"`
}

var Customers = []Customer{
	{Name: "Arief", City: "Jakarta", Zipcode: "12510"},
	{Name: "Darmawan", City: "Bogor", Zipcode: "16969"},
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(Customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Customers)
	}
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(vars)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vars)
	}
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {

}
