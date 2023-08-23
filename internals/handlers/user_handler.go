package handlers

import (
	"github.com/gofiber/fiber/v2"
	"newglo/internals/domain"
	"newglo/internals/ports"
)

type UserHandlers struct {
	userService ports.UserService
}

//	func NewUserService(repository ports.UserRepository) *UserService {
//		return &UserService{
//			UserRepository: repository,
//		}
//	}
type loginGetter interface {
	Login(user domain.LoginCredentials) error
}
type Auth struct {
	loginGetter loginGetter
}

func (db Auth) Login(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewApp(foo loginGetter) *Auth {
	return &Auth{foo}
}
func (h *UserHandlers) Login(c *fiber.Ctx) error {
	var email string
	var password string
	//Extract the body and get the email and password
	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}
