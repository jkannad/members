package models


type Member struct {
	Id            int `json:"id" isRequired:"false" length:"0"`
	Title         string `json:"title" isRequired:"true" length:"5"`
	FirstName     string `json:"first_name" isRequired:"true" length:"30"`
	LastName      string `json:"last_name" isRequired:"true" length:"30"`
	Gender        string `json:"gender" isRequired:"true" length:"6"`
	Dob           string `json:"dob" isRequired:"false" length:"10"`
	Doj           string `json:"doj" isRequired:"true" length:"10"`
	Address1      string `json:"address1" isRequired:"true" length:"50"`
	Address2      string `json:"address2" isRequired:"false" length:"50"`
	Area		  string `json:"area" isRequired:"false" length:"50"`
	Country       string `json:"country" isRequired:"true" length:"50"`
	State         string `json:"state" isRequired:"true" length:"50"`
	City          string `json:"city" isRequired:"false" length:"50"`
	PostalCode    string `json:"postal_code" isRequired:"true" length:"10"`
	ContactNumber string `json:"contact_number" isRequired:"true" length:"15"`
	DialCode	  string `json:"dial_code" isRequired:"false"`
	Email         string `json:"email" isRequired:"false" length:"150"`
	CreatedAt	  string `json:"created_at" isRequired:"false"`
	UpdatedAt     string `json:"updated_at" isRequired:"false"`
	CreatedBy     string `json:"created_by" isRequired:"false"`
	UpdatedBy     string `json:"updated_by" isRequired:"false"`
}



