package service

import (
	"os"
	"io"
	"encoding/json"
	"github.com/jkannad/spas/members/internal/models"
	"github.com/jkannad/spas/members/internal/repo"
	"github.com/jkannad/spas/members/internal/helper"
)

const (
	COUNTRY_DATA_FILE = "data/countries.json"
	STATE_DATA_FILE = "data/states.json"
	CITY_DATA_FILE = "data/cities.json"
	DIAL_CODE_FILE = "data/dialcodes.json"
	
)	

//LoadCountries loads countries details to database
func LoadCountries() error {
	file, err := os.Open(COUNTRY_DATA_FILE)
	if err != nil {
		helper.ServiceError(err)
		return err
	}
	defer file.Close()

	byteValues, err := io.ReadAll(file)
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	var countries []models.Country

	err = json.Unmarshal(byteValues, &countries)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	//call repo to insert data to database

	err = repo.LoadCountries(countries)

	if err != nil {
		return err
	}
	return nil
}

//LoadStates loads states details to database
func LoadStates() error {
	file, err := os.Open(STATE_DATA_FILE)
	if err != nil {
		helper.ServiceError(err)
		return err
	}
	defer file.Close()

	byteValues, err := io.ReadAll(file)
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	var countries map[string][]models.State

	err = json.Unmarshal(byteValues, &countries)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	var allStates []models.State

	for _, states := range countries {
		allStates = append(allStates, states...)
	}

	//call repo to insert data to database
	err = repo.LoadStates(allStates)

	if err != nil {
		return err
	}

	return nil
}

//LoadCities loads cities details to database
func LoadCities() error {
	file, err := os.Open(CITY_DATA_FILE)
	if err != nil {
		helper.ServiceError(err)
		return err
	}
	defer file.Close()

	byteValues, err := io.ReadAll(file)
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	var counties map[string]map[string][]models.City

	err = json.Unmarshal(byteValues, &counties)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	var cities []models.City

	for _, states := range counties {
		for _, cty := range states {
			cities = append(cities, cty...)
		}
	}

	err = repo.LoadCities(cities)

	if err != nil {
		return err
	}

	return nil
}

func LoadDialCodes() error {
	file, err := os.Open(DIAL_CODE_FILE)
	if err != nil {
		helper.ServiceError(err)
		return err
	}
	defer file.Close()

	byteValues, err := io.ReadAll(file)
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	var dialCodes []models.DialCode

	err = json.Unmarshal(byteValues, &dialCodes)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	err = repo.LoadDialCodes(dialCodes)

	if err != nil {
		return err
	}
	return nil
}

//GetCountries returns all the countries details
func GetCountries() ([]models.Country, error) {
	countries, err := repo.GetCountries()
	if err != nil {
		return nil, err
	}
	return countries, nil
}

//GetStates returns all the states for the given country code
func GetStates(countryCD string) ([]models.State, error) {
	states, err := repo.GetStates(countryCD)
	if err != nil {
		return nil, err
	}
	return states, nil
}

//GetCitites returns all the cities of the selected state and country
func GetCities(countryCD, stateCD string) ([]models.City, error) {
	cities, err := repo.GetCities(countryCD, stateCD)
	if err != nil {
		return nil, err
	}
	return cities, nil
}

//GetDialCodes returns the dial codes of all the countries
func GetDialCodes() (map[string]models.DialCode, error) {
	dialCodes, err := repo.GetDialCodes()
	if err != nil {
		return nil, err
	}
	return dialCodes, nil
}

//GetDialCodes returns the dial codes of all the countries
func GetDialCode(countryCode string) (models.DialCode, error) {
	dialCode, err := repo.GetDialCode(countryCode)
	if err != nil {
		return models.DialCode{}, err
	}
	return dialCode, nil
}
