package models

type Member struct {
	Id            int    `json:"id" isRequired:"true" length:"30"`
	Title         string `json:"title" isRequired:"true" length:"3"`
	FirstName     string `json:"first_name" isRequired:"true" length:"30"`
	LastName      string `json:"last_name" isRequired:"true" length:"30"`
	Sex           string `json:"sex" isRequired:"true" length:"6"`
	Dob           string `json:"dob" isRequired:"true" length:"10"`
	Doj           string `json:"doj" isRequired:"true" length:"10"`
	Address1      string `json:"address1" isRequired:"true" length:"50"`
	Address2      string `json:"address2" isRequired:"true" length:"50"`
	Country       string `json:"country" isRequired:"true" length:"50"`
	State         string `json:"state" isRequired:"true" length:"50"`
	City          string `json:"city" isRequired:"true" length:"50"`
	PostalCode    string `json:"postal_code" isRequired:"true" length:"15"`
	ContactNumber string `json:"contact_number" isRequired:"true" length:"15"`
	Email         string `json:"email" isRequired:"true" length:"150"`
}