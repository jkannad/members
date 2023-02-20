package repo

import (
	"fmt"
	"log"
	"database/sql"
	"strings"
	"github.com/jkannad/spas/members/internal/models"
	"github.com/jkannad/spas/members/internal/helper"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_ENGINE = "sqlite3"
	DB_NAME = "data/spaa.db"
	INSERT_MEMBER_SQL = `insert into member(title, first_name, last_name, gender, dob, doj, address1,
							address2, area, country, state, city, postal_code, dial_code, phone, email) 
						   	values( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	UPDATE_MEMBER_SQL = `update member set title=?, first_name=?, last_name=?, gender=?, dob=?,
						   	doj=?, address1=?, address2=?, area=?, country=?, state=?, city=?, postal_code=?, dial_code=?, phone=?, email=?
						   	where id=?`
	SELECT_MEMBER_BY_ID_SQL = `select id, title, first_name, last_name, gender, dob, doj, address1,
								 	address2, area, country, state, city, postal_code, dial_code, phone, email 
								 	from member where id=?`
	SELECT_MEMBER_ALL = `select id, first_name, last_name, gender, area, cn.name country, st.name state, city, postal_code, dial_code, phone, email 
							from member m, countries cn, states st
							where m.country=cn.cd and m.state=st.cd and cn.cd = st.country_cd order by first_name asc`
	SEARCH_MEMBERS = `select id, first_name, last_name, gender, area, cn.name country, st.name state, city, postal_code, dial_code, phone, email 
						from member m, countries cn, states st
						where m.country=cn.cd and m.state=st.cd and cn.cd = st.country_cd `
)

//AddMember add a member to the database 
func AddMember(member *models.Member) error {
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare(INSERT_MEMBER_SQL)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		member.Title,
		member.FirstName,
		member.LastName,
		member.Gender,
		member.Dob,
		member.Doj,
		member.Address1,
		member.Address2,
		member.Area,
		member.Country,
		member.State,
		member.City,
		member.PostalCode,
		member.DialCode,
		member.ContactNumber,
		member.Email,
	)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	log.Printf("Last inserted Id %d. Affected rows %d\n", lastInsertId, rowsAffected)
	return nil
}

//UpdateMember updates a member to the database
func UpdateMember(member *models.Member) error{
	db, err := sql.Open(DB_ENGINE, DB_NAME)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	stm, err := db.Prepare(UPDATE_MEMBER_SQL)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stm.Close()

	result, err := stm.Exec(
		member.Title,
		member.FirstName,
		member.LastName,
		member.Gender,
		member.Dob,
		member.Doj,
		member.Address1,
		member.Address2,
		member.Area,
		member.Country,
		member.State,
		member.City,
		member.PostalCode,
		member.DialCode,
		member.ContactNumber,
		member.Email,
		member.Id,
	)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	log.Printf("Last updated Id %d. Affected rows %d\n", lastInsertId, rowsAffected)
	return nil
}

//GetMemberById get a member by Id
func GetMemberById(id int) (*models.Member, error){
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return &models.Member{}, err
	}

	defer db.Close()

	stmt, err := db.Prepare(SELECT_MEMBER_BY_ID_SQL)

	if err != nil {
		helper.ServiceError(err)
		return &models.Member{}, err
	}

	defer stmt.Close()

	var member models.Member
	err = stmt.QueryRow(id).Scan(&member.Id, 
								  &member.Title,
								  &member.FirstName,
								  &member.LastName,
								  &member.Gender,
								  &member.Dob,
								  &member.Doj,
								  &member.Address1,
								  &member.Address2,
								  &member.Area,
								  &member.Country,
								  &member.State,
								  &member.City,
								  &member.PostalCode,
								  &member.DialCode,
								  &member.ContactNumber,
								  &member.Email,
								)
	
	if err != nil {
		helper.ServiceError(err)
		return &models.Member{}, err
	}
	return &member, nil
}

//GetMembers get slice of members from database based on matching conditions
func GetMembers(member *models.Member) ([]models.Member, error){
	return []models.Member{}, nil
}

//GetAllMembers retrieves all the members details from database
func GetAllMembers() ([]models.Member, error){
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return []models.Member{}, err
	}

	defer db.Close()

	rows, err := db.Query(SELECT_MEMBER_ALL)

	if err != nil {
		helper.ServiceError(err)
		return []models.Member{}, err
	}

	var members []models.Member

	for rows.Next() {
		var member models.Member
		err = rows.Scan(&member.Id, 
			&member.FirstName,
			&member.LastName,
			&member.Gender,
			&member.Area,
			&member.Country,
			&member.State,
			&member.City,
			&member.PostalCode,
			&member.DialCode,
			&member.ContactNumber,
			&member.Email,
		  )
		  if err != nil {
			helper.ServiceError(err)
			return []models.Member{}, err
		  }

		  members = append(members, member)
	}	
	
	return members, nil
	
}

func SearchMembers(search models.Search) ([]models.Member, error) {
	var condition string
	if cn := strings.Trim(search.ContactNumber, ""); len(cn) != 0 {
		condition = fmt.Sprintf("and phone = '%s'", cn)
		return executeSearchQuery(condition)
	}

	if email := strings.Trim(search.Email, ""); len(email) != 0 {
		condition = fmt.Sprintf("and email = '%s'", email)
		return executeSearchQuery(condition)
	}

	if fname := strings.Trim(search.FirstName, ""); len(fname) != 0 {
		condition = fmt.Sprintf("and first_name like '%s%%'", fname)
		return executeSearchQuery(condition)
	}

	if lname := strings.Trim(search.LastName, ""); len(lname) != 0 {
		condition = fmt.Sprintf("and last_name like '%s%%'", lname)
		return executeSearchQuery(condition)
	}

	if area := strings.Trim(search.Area, ""); len(area) != 0 {
		condition = fmt.Sprintf("and area like '%%%s%%'", area)
		return executeSearchQuery(condition)
	}

	if postalCode := strings.Trim(search.PostalCode, ""); len(postalCode) != 0 {
		condition = fmt.Sprintf("and postal_code like '%s%%'", postalCode)
		return executeSearchQuery(condition)
	}

	if len(search.Country) != 0 && len(search.State) !=0 && len(search.City) !=0 {
		condition = fmt.Sprintf("and country = '%s' and state = '%s' and city = '%s'", search.Country, search.State, search.City)
		return executeSearchQuery(condition)
	}

	if len(search.Country) != 0 && len(search.State) !=0 {
		condition = fmt.Sprintf("and country = '%s' and state = '%s'", search.Country, search.State)
		return executeSearchQuery(condition)
	}

	if len(search.Country) != 0 {
		condition = fmt.Sprintf("and country = '%s'", search.Country)
		return executeSearchQuery(condition)
	}

	return GetAllMembers()
}

func executeSearchQuery(condition string) ([]models.Member, error) {
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return []models.Member{}, err
	}

	defer db.Close()

	rows, err := db.Query(SEARCH_MEMBERS + condition)

	if err != nil {
		helper.ServiceError(err)
		return []models.Member{}, err
	}

	var members []models.Member

	for rows.Next() {
		var member models.Member
		err = rows.Scan(&member.Id, 
			&member.FirstName,
			&member.LastName,
			&member.Gender,
			&member.Area,
			&member.Country,
			&member.State,
			&member.City,
			&member.PostalCode,
			&member.DialCode,
			&member.ContactNumber,
			&member.Email,
		  )
		  if err != nil {
			helper.ServiceError(err)
			return []models.Member{}, err
		  }

		  members = append(members, member)
	}	
	
	return members, nil
}

