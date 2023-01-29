package models

import (
	"reflect"
)

type Member struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Sex           string `json:"sex"`
	Dob           string `json:"dob"`
	Doj           string `json:"doj"`
	Address1      string `json:"address1"`
	Address2      string `json:"address2"`
	Country       string `json:"country"`
	State         string `json:"state"`
	City          string `json:"city"`
	PostalCode    string `json:"postal_code"`
	ContactNumber string `json:"contact_number"`
	Email         string `json:"email"`
}

type ValidationConfig struct {
	IsRequired bool
	Length     int
}

type MemberValidator struct {
	Id            int    `name:"id" isRequired:"true" length:"30"`
	Title         string `name:"title" isRequired:"true" length:"3"`
	FirstName     string `name:"first_name" isRequired:"true" length:"30"`
	LastName      string `name:"last_name" isRequired:"true" length:"30"`
	Sex           string `name:"sex" isRequired:"true" length:"6"`
	Dob           string `name:"dob" isRequired:"true" length:"10"`
	Doj           string `name:"doj" isRequired:"true" length:"10"`
	Address1      string `name:"address1" isRequired:"true" length:"50"`
	Address2      string `name:"address2" isRequired:"true" length:"50"`
	Country       string `name:"country" isRequired:"true" length:"50"`
	State         string `name:"state" isRequired:"true" length:"50"`
	City          string `name:"city" isRequired:"true" length:"50"`
	PostalCode    string `name:"postal_code" isRequired:"true" length:"15"`
	ContactNumber string `name:"contact_number" isRequired:"true" length:"15"`
	Email         string `name:"email" isRequired:"true" length:"150"`
}

func (m *MemberValidator) GetValidationConfig(field string) (ValidationConfig, error) {
	typ := reflect.TypeOf(field)

	f, _ := typ.FieldByName(field)

}
