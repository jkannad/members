package models

type Search struct {
	FirstName string 		`json:"first_name"`
	LastName string			`json:"last_name"`
	Area string				`json:"area"`
	PostalCode string		`json:"postal_code"`
	Country string			`json:"country"`
	State string			`json:"state"`
	City string				`json:"city"`
	ContactNumber string 	`json:"contact_number"`
	Email string			`json:"email"`
}