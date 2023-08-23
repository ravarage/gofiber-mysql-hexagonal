package services

import "newglo/internals/ports"

type UserService struct {
	userRepository ports.UserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.UserService = (*UserService)(nil)

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(email string, password string) error {

	return nil
}
