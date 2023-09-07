package service

import (
	BIP_project "github.com/nekitalek/bip_project/backend"
	"github.com/nekitalek/bip_project/backend/internal/repository"
)

type UserService struct {
	repo repository.Authorization
}

func NewUserService(repo repository.Authorization) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(user_id int) (BIP_project.User_data, error) {
	var user_data BIP_project.User_data
	user_auth, err := s.repo.GetUserById(user_id)
	if err != nil {
		return user_data, err
	}
	user_data.User_id = user_auth.User_id
	user_data.Login = user_auth.Login
	user_data.Username = user_auth.Username

	return user_data, nil
}

// func (s *UserService) UpdateUser(user_id int) (BIP_project.User, error) {

// }
