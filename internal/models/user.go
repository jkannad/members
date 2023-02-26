package models

type User struct {
	UserName 		string `json:"user_name"`
	FirstName 		string `json:"first_name"`
	LastName 		string `json:"last_name"`
	DialCode		string `json:"dial_code"`
	ContactNumber 	string `json:"contact_number"`
	Email 			string `json:"email"`
	Password 		string `json:"password"`
	AccessLevel     int	   `json:"access_lavel"`
	CreatedAt	  	string `json:"created_at"`
	UpdatedAt     	string `json:"updated_at"`
	CreatedBy     	string `json:"created_by"`
	UpdatedBy     	string `json:"updated_by"`
}