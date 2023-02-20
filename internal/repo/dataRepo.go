package repo

import (
	"database/sql"
	"strings"
	"log"
	"github.com/jkannad/spas/members/internal/models"
	"github.com/jkannad/spas/members/internal/helper"
)

const (
	INSERT_COUNTRY = "insert into countries (cd, name) values "
	INSERT_STATE = "insert into states (cd, country_cd, name) values "
	INSERT_CITY = "insert into cities (name, state_cd, country_cd) values "
	INSERT_DIAL_CODES = "insert into dial_codes (country_cd, dial_cd) values "
	SELECT_COUNTRIES = "select * from countries order by name"
	SELECT_STATES = "select cd, country_cd, name from states where country_cd=? order by name asc"
	SELECT_CITIES = "select country_cd, state_cd, name from cities where country_cd=? and state_cd=? order by name asc"
	SELECT_DIAL_CODES = "select country_cd, dial_cd from dial_codes"
	SELECT_DIAL_CODE = "select country_cd, dial_cd from dial_codes where country_cd = ?"
)

//LoadCountries loads all the countries details to DB
func LoadCountries(countries []models.Country) error {
	var inserts []string
    var params []interface{}
    for _, country := range countries {
        inserts = append(inserts, "(?, ?)")
        params = append(params, country.IsoCode, country.Name)
    }
    queryVals := strings.Join(inserts, ",")
    query := INSERT_COUNTRY + queryVals

	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	log.Println("Counties were loaded")
	log.Printf("Last inserted Id %d. Affected rows %d\n", lastInsertId, rowsAffected)

	return nil
}

//LoadStates loads all the states details to DB
func LoadStates(states []models.State) error {
	var inserts []string
    var params []interface{}
    for _, state := range states {
        inserts = append(inserts, "(?, ?, ?)")
        params = append(params, state.IsoCode, state.CountryCode, state.Name)
    }
    queryVals := strings.Join(inserts, ",")
    query := INSERT_STATE + queryVals

	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	log.Println("States were loaded")
	log.Printf("Last inserted Id %d. Affected rows %d\n", lastInsertId, rowsAffected)

	return nil
}

//LoadCities loads all the cities details to DB
func LoadCities(cities []models.City) error {
	var inserts []string
    var params []interface{}

    for i, city := range cities {
        inserts = append(inserts, "(?, ?, ?)")
        params = append(params, city.Name, city.StateCode, city.CountryCode)
		if i % 10000 == 0 {
			err := insertCities(inserts, params)
			if err != nil {
				helper.ServiceError(err)
				return err
			}
			inserts = nil
			params = nil
		}
    }
    
	if len(inserts) > 0 && len(params) > 0 {
		err := insertCities(inserts, params)
		if err != nil {
			helper.ServiceError(err)
			return err
		}
	}

	return nil
}

func insertCities(inserts []string, params []interface{}) error {
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	queryVals := strings.Join(inserts, ",")
   	query := INSERT_CITY + queryVals

	stmt, err := db.Prepare(query)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stmt.Close()

	

	result, err := stmt.Exec(params...)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	log.Println("Cities were loaded")
	log.Printf("Last inserted Id %d. Affected rows %d\n", lastInsertId, rowsAffected)

	return nil
}

//LoadCountries loads all the countries details to DB
func LoadDialCodes(dialCodes []models.DialCode) error {
	var inserts []string
    var params []interface{}
    for _, dialCode := range dialCodes {
        inserts = append(inserts, "(?, ?)")
        params = append(params, dialCode.CountryCD, dialCode.DialCD)
    }
    queryVals := strings.Join(inserts, ",")
    query := INSERT_DIAL_CODES + queryVals

	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(params...)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	log.Println("Dial codes were loaded")
	log.Printf("Last inserted Id %d. Affected rows %d\n", lastInsertId, rowsAffected)

	return nil
}

func GetCountries() ([]models.Country, error){
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query(SELECT_COUNTRIES)
	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	var countries []models.Country
	
	for rows.Next() {
		var country models.Country
		err = rows.Scan(&country.IsoCode, &country.Name)
		if err != nil {
			helper.ServiceError(err)
			return nil, err
		}
		countries = append(countries, country)
	}

	return countries, nil
}

func GetStates(countryCode string) ([]models.State, error){
	db, err := sql.Open(DB_ENGINE, DB_NAME)

	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	defer db.Close()

	stm, err := db.Prepare(SELECT_STATES)

	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	defer stm.Close()

	var states []models.State

	rows, err := stm.Query(countryCode)
	
	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	for rows.Next() {
		var state models.State
		err = rows.Scan(&state.IsoCode,
						&state.CountryCode,
						&state.Name,
						)
		if err != nil {
			helper.ServiceError(err)
			return nil, err
		}
		states = append(states, state)
	}
	return states, nil
}

func GetCities(countryCode string, stateCode string) ([]models.City, error){
	db, err := sql.Open(DB_ENGINE, DB_NAME)

	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	defer db.Close()

	stm, err := db.Prepare(SELECT_CITIES)

	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	defer stm.Close()

	var cities []models.City

	rows, err := stm.Query(countryCode, stateCode)
	
	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	for rows.Next() {
		var city models.City
		err = rows.Scan(&city.CountryCode, &city.StateCode, &city.Name)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, nil
}

func GetDialCodes() (map[string]models.DialCode, error){
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query(SELECT_DIAL_CODES)
	if err != nil {
		helper.ServiceError(err)
		return nil, err
	}

	dialCodes := make(map[string]models.DialCode)
	
	for rows.Next() {
		var dialCode models.DialCode
		err = rows.Scan(&dialCode.CountryCD, &dialCode.DialCD)
		if err != nil {
			helper.ServiceError(err)
			return nil, err
		}
		dialCodes[dialCode.CountryCD] = dialCode
	}

	return dialCodes, nil
}

func GetDialCode(countryCode string) (models.DialCode, error) {
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return models.DialCode{}, err
	}

	defer db.Close()

	stm, err := db.Prepare(SELECT_DIAL_CODE)

	if err != nil {
		helper.ServiceError(err)
		return models.DialCode{}, err
	}

	defer stm.Close()

	var dialCode models.DialCode

	stm.QueryRow(countryCode).Scan(&dialCode.CountryCD, &dialCode.DialCD)

	return dialCode, nil
	
}

