package domain

import (
	"github.com/rokafela/udemy-banking/dto"
	"github.com/rokafela/udemy-banking/errs"
)

// domain
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string ``
	City        string ``
	Zipcode     string ``
	DateOfBirth string `db:"date_of_birth"`
	Status      string ``
}

// secondary port
type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}

func (c Customer) statusAsText() string {
	customerStatus := "active"
	if c.Status == "0" {
		customerStatus = "inactive"
	}
	return customerStatus
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}
