package config

import (
	"reflect"
	"strconv"
	"github.com/jkannad/spas/members/internal/models"

)

var FormNames = []string{"member"}

type FieldConfig struct {
	IsRequired bool
	Length     int
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

func BuildFormFieldConfigs(formFieldConfigs map[string]map[string]FieldConfig){
	for _, formName := range FormNames {
		formFieldConfigs[formName] = GetFormFieldConfigs(models.Member{})
	}

} 

