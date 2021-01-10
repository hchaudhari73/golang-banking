package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"101", "Harshal", "Dombivli", "421202", "1999-30-02", "1"},
		{"102", "Bruce", "Goutam", "6770058", "1994-30-02", "1"},
	}
	return CustomerRepositoryStub{customers}
}
