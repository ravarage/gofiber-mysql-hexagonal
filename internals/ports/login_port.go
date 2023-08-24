package ports

import (
	"github.com/gofiber/fiber/v2"
	"newglo/internals/domain"
)

type UserService interface {
	Login(userCredentials domain.LoginCredentials) (int, error)
}

type UserRepository interface {
	DoLogin(user domain.LoginCredentials) (int, error)
}

type UserHandlers interface {
	Login(c *fiber.Ctx) error
}

type UserMiddlewares interface {
	Login(c *fiber.Ctx) error
}
