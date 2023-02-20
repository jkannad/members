package config

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"github.com/jkannad/spas/members/internal/models"

)

var FormNames = []string{"member"}

type FieldConfig struct {
	IsRequired bool
	Length     int
}

func BuildFormFieldConfigs() map[string]map[string]FieldConfig {
	var formFieldConfigs = make(map[string]map[string]FieldConfig)
	for _, formName := range FormNames {
		formFieldConfigs[formName] = getFieldConfigsByFrom(models.Member{})
	}
	return formFieldConfigs
}

func getFieldConfigsByFrom(model interface{}) map[string]FieldConfig {
	fvc := make(map[string]FieldConfig)
	typ := reflect.TypeOf(model)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		isRequired, _ := strconv.ParseBool(f.Tag.Get("isRequired"))
		length, _ := strconv.Atoi(f.Tag.Get("length"))

		fvc[f.Name] = FieldConfig {
			IsRequired: isRequired,
			Length: length,
		}
	}
	return fvc
}

func ValidateFormData(formData interface{}, fieldConfig map[string]FieldConfig) map[string][]string {
	error := make(map[string][]string)
	element := reflect.ValueOf(formData).Elem()
	elementType := element.Type()
	for i := 0; i < element.NumField(); i++ {
		field := elementType.Field(i)
		name := field.Name
		value := element.Field(i).Interface()
		fc, found := fieldConfig[name]
		if !found {
			log.Fatalf("Element %s is not found in the form field config map\n", name)
		} else {
			isRequired := fc.IsRequired
			length := fc.Length
			val := fmt.Sprintf("%v", value)
			if isRequired {
				if _, ok := value.(string); ok {
					if len(strings.TrimSpace(val)) == 0 {
						error[name] = append(error[name], fmt.Sprintf("Field '%s' cannot have empty value", name))
					}
				}
				if len(val) > length {
					error[name] = append(error[name], fmt.Sprintf("Length of field '%s' cannot be greater than '%d'", name, length))
				}
			}
		}
	}
	return error
}