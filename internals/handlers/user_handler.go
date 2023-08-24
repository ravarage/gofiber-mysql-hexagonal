package handlers

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"newglo/internals/domain"
	"newglo/internals/ports"
)

type UserHandlers struct {
	userService ports.UserService
}

func NewApp(foo ports.UserService) *UserHandlers {

	return &UserHandlers{userService: foo}
}
func (h *UserHandlers) Login(c *fiber.Ctx) error {

	var userCredentials domain.LoginCredentials
	//use the json.Unmarshal to extract the body
	err := json.Unmarshal(c.Body(), &userCredentials)
	if err != nil {

		return err
	}

	//Extract the body and get the email and password
	a, err := h.userService.Login(userCredentials)

	if err != nil {
		fmt.Print(err)
		return err
	}
	//store a in secure cookie for infinite time
	c.Cookie(&fiber.Cookie{
		Name:     "a",
		Value:    fmt.Sprintf("%d", a),
		HTTPOnly: true,
	})

	return nil
}
