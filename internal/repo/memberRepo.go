package repo

import (
	"fmt"
	"log"
	"database/sql"
	"strings"
	"time"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/jkannad/spas/members/internal/models"
	"github.com/jkannad/spas/members/internal/helper"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_ENGINE = "sqlite3"
	DB_NAME = "data/spaa.db"
	INSERT_MEMBER_SQL = `insert into member(title, first_name, last_name, gender, dob, doj, address1,
							address2, area, country, state, city, postal_code, dial_code, phone, email, created_by, created_at) 
						   	values( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	UPDATE_MEMBER_SQL = `update member set title=?, first_name=?, last_name=?, gender=?, dob=?,
						   	doj=?, address1=?, address2=?, area=?, country=?, state=?, city=?, postal_code=?, dial_code=?, phone=?, email=?, updated_by=?, updated_at=?
						   	where id=?`
	SELECT_MEMBER_BY_ID_SQL = `select id, title, first_name, last_name, gender, dob, doj, address1,
								 	address2, area, country, state, city, postal_code, dial_code, phone, email 
								 	from member where id=?`
	SELECT_MEMBER_ALL_SQL = `select id, first_name, last_name, gender, area, cn.name country, st.name state, city, postal_code, dial_code, phone, email 
								from member m, countries cn, states st
							where m.country=cn.cd and m.state=st.cd and cn.cd = st.country_cd order by first_name asc`
	SEARCH_MEMBERS_SQL = `select id, first_name, last_name, gender, area, cn.name country, st.name state, city, postal_code, dial_code, phone, email 
							from member m, countries cn, states st
							where m.country=cn.cd and m.state=st.cd and cn.cd = st.country_cd `
	INSERT_USER_SQL = `insert into user (user_name, first_name, last_name, password, dial_code, phone, email, access_level, created_by, created_at)
						values (?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
	UPDATE_PASSWORD_SQL = `update user set password=?, updated_by=?, updated_at=? where user_name=?`
	UPDATE_USER_SQL = `update user set first_name=?, last_name=?, dial_code=?, phone=?, email=?, updated_by=?, updated_at=? where user_name=?`
	SELECT_USER_BY_NAME_SQL = `select user_name, first_name, last_name, phone, email, access_level from user where user_name = ?`
	VALIDATE_USER_SQL = `select password from user where user_name=?`
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

	currentTime := time.Now()

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
		member.CreatedBy,
		currentTime.String(),
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
		member.UpdatedBy,
		time.Now().String(),
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

//AddUser adds the user
func AddUser(user *models.User) error {
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare(INSERT_USER_SQL)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		user.UserName,
		user.FirstName,
		user.LastName,
		user.DialCode,
		user.ContactNumber,
		user.Email,
		user.AccessLevel,
		user.CreatedBy,
		time.Now().String(),
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

//UpdateUser updates the user details
func UpdateUser(user *models.User) error {
	db, err := sql.Open(DB_ENGINE, DB_NAME)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	stm, err := db.Prepare(UPDATE_USER_SQL)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stm.Close()

	result, err := stm.Exec(
		user.FirstName,
		user.LastName,
		user.DialCode,
		user.ContactNumber,
		user.Email,
		user.UpdatedBy,
		time.Now().String(),
		user.UserName,
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

//UpdatePassword updates the user's password
func UpdatePassword(user *models.User) error {
	db, err := sql.Open(DB_ENGINE, DB_NAME)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	stm, err := db.Prepare(UPDATE_PASSWORD_SQL)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stm.Close()

	result, err := stm.Exec(
		user.Password,
		user.UpdatedBy,
		time.Now().String(),
		user.UserName,
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
func GetUserByName(userName string) (*models.User, error){
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return &models.User{}, err
	}

	defer db.Close()

	stmt, err := db.Prepare(SELECT_USER_BY_NAME_SQL)

	if err != nil {
		helper.ServiceError(err)
		return &models.User{}, err
	}

	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(userName).Scan(&user.UserName, 
								  &user.FirstName,
								  &user.LastName,
								  &user.ContactNumber,
								  &user.Email,
								)
	
	if err != nil {
		helper.ServiceError(err)
		return &models.User{}, err
	}
	return &user, nil
}

//AuthenticateUser authenticates a user in the database
func AuthenticateUser(userName, password string) error{
	db, err := sql.Open(DB_ENGINE, DB_NAME)
	
	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare(VALIDATE_USER_SQL)

	if err != nil {
		helper.ServiceError(err)
		return err
	}

	defer stmt.Close()

	var hashedPassword string
	err = stmt.QueryRow(userName).Scan(hashedPassword, 
								  userName,
								)
	if err != nil {
		helper.ServiceError(err)
		return err
	} 

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("incorrect password")
	} else if err != nil {
		return err
	}
	return nil
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

	rows, err := db.Query(SELECT_MEMBER_ALL_SQL)

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

	rows, err := db.Query(SEARCH_MEMBERS_SQL + condition)

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

