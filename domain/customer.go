package domain

import "github.com/rokafela/udemy-banking/errs"

// domain
type Customer struct {
	Id          string `json:"customer_id" xml:"customerId" db:"customer_id"`
	Name        string `json:"full_name" xml:"fullName"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zipCode"`
	DateOfBirth string `json:"date_of_birth" xml:"dateOfBirth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

// secondary port
type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
