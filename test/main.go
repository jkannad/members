package main

import (
	"strconv"
	"fmt"
	"reflect"
)

type Member struct {
	Id            int    `json:"id" isRequired:"true" length:"30"`
	Title         string `json:"title" isRequired:"true" length:"3"`
	FirstName     string `json:"first_name" isRequired:"true" length:"30"`
	LastName      string `json:"last_name" isRequired:"true" length:"30"`
	Sex           string `json:"sex" isRequired:"true" length:"6"`
}

var FormNames = []string{"member"}

type FieldConfig struct {
	IsRequired bool
	Length     int
}

func main() {
	var formFieldConfigs = make(map[string]map[string]FieldConfig)
	BuildFormFieldConfigs(formFieldConfigs)
	for k, v := range formFieldConfigs {
		fmt.Printf("Values for the form '%s' are given below\n", k)
		for k1, v1 := range v {
			fmt.Printf("Key %s value {IsRequired: %v, Length: %d}\n", k1, v1.IsRequired, v1.Length)
		}
	}
}

func BuildFormFieldConfigs(formFieldConfigs map[string]map[string]FieldConfig){
	for _, formName := range FormNames {
		formFieldConfigs[formName] = GetFormFieldConfigs(Member{})
	}
} 

func GetFormFieldConfigs(model interface{}) map[string]FieldConfig {
	fvc := make(map[string]FieldConfig)
	typ := reflect.TypeOf(model)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		isRequired, _ := strconv.ParseBool(f.Tag.Get("isRequired"))
		length, _ := strconv.Atoi(f.Tag.Get("length"))
		
		fvc[f.Tag.Get("json")] = FieldConfig {
			IsRequired: isRequired,
			Length: length,
		}
	}
	return fvc
}
/*
func default_tag(p Person) string {

	// TypeOf returns type of
	// interface value passed to it
	typ := reflect.TypeOf(p)
	fmt.Printf("Type of p %v\n", typ)

	// checking if null string
	if p.name == "" {

		// returns the struct field
		// with the given parameter "name"

		f, _ := typ.FieldByName("name")

		// returns the value associated
		// with key in the tag string
		// and returns empty string if
		// no such key in tag
		p.name = f.Tag.Get("default2")
	}

	return fmt.Sprintf("%s", p.name)
}*/

