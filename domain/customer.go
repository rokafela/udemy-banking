package domain

// domain
type Customer struct {
	Id          string `json:"customer_id" xml:"customerId"`
	Name        string `json:"full_name" xml:"fullName"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zipCode"`
	DateOfBirth string `json:"date_of_birth" xml:"dateOfBirth"`
	Status      string `json:"status" xml:"status"`
}

// secondary port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
