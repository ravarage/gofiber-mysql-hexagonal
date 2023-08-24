package services

import (
	"fmt"
	"newglo/internals/domain"
	"newglo/internals/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
//var _ ports.UserService = (*UserService)(nil)

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(credentials domain.LoginCredentials) (int, error) {
	fmt.Println("Login service")
	a, err := s.userRepository.DoLogin(credentials)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return a, nil
}
