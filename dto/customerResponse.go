package dto

/*
	We we're directly returning the domain object to the user,
	which was exposing the domain to the user,
	so we have to instead return a response
	rather than passing the object from the domain.
	That is why we are creating dto: which will be passed as response to the user
*/

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}
