package service

import (
	"github.com/jkannad/spas/members/internal/models"
	"github.com/jkannad/spas/members/internal/repo"
	"errors"
)


//Upsert insert or update a member to the database
func Upsert(m *models.Member) error {
	if m.Id == 0 {
		err := repo.AddMember(m)
		if err != nil {
			return err
		}
	} else {
		err := repo.UpdateMember(m)
		if err != nil {
			return err
		}
	}
	return nil
}

//GetMember get member details by passing id
func GetMember(id int) (*models.Member, error) {
	if id == 0 {
		return &models.Member{}, errors.New("member id cannot be zero")
	}

	member, err := repo.GetMemberById(id)

	if err != nil {
		return &models.Member{}, err
	}

	return member, nil
} 

//GetAllMembers get all the members' details
func GetAllMembers() ([]models.Member, error) {
	var members []models.Member

	members, err := repo.GetAllMembers()
	
	if err != nil {
		return []models.Member{}, err
	}
	return members, nil
}

//SearchMembers return matching members with search criteria
func SearchMembers(search models.Search) ([]models.Member, error) {
	var members []models.Member

	members, err := repo.SearchMembers(search)
	
	if err != nil {
		return []models.Member{}, err
	}
	return members, nil
}

