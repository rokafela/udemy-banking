package domain

// adapter for secondary port
type RandomStub struct {
	customers []Customer
}

func (s RandomStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewRandomStub() RandomStub {
	customers := []Customer{
		{Id: "1", Name: "Arief", City: "Jakarta", Zipcode: "12510", DateOfBirth: "1990-01-01", Status: "1"},
		{Id: "2", Name: "Darmawan", City: "Bogor", Zipcode: "16969", DateOfBirth: "2001-02-02", Status: "1"},
	}
	return RandomStub{customers}
}
