package server

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/utils"
	jwtware "github.com/gofiber/jwt/v3"
	"newglo/internals/handlers"
	"newglo/internals/ports"
	"time"
)

const (
	HeaderName = "X-Csrf-Token"
)

type Server struct {
	//We will add every new Handler here
	userHandlers ports.UserHandlers
	//middlewares ports.IMiddlewares
	//paymentHandlers ports.IPaymentHandlers
}

func NewServer(uHandlers *handlers.UserHandlers) *Server {
	return &Server{
		userHandlers: uHandlers,
		//paymentHandlers: pHandlers
	}
}

func (s *Server) Initialize() {

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	v1 := app.Group("/api")

	userRoutes := v1.Group("/user")
	userRoutes.Post("/login", s.userHandlers.Login)
	app.Get("/metrics", monitor.New())

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "http://localhost:5173,https://gloctor.com,https://cms.gloctor.com,http://gloctor.com,http://cms.gloctor.com,https://api.gloctor.com",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: true,
		ExposeHeaders:    "",
	}))
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:" + HeaderName,
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ErrorHandler:   defaultErrorHandler,
		Extractor:      CsrfFromHeader(HeaderName),
	}))
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "secret-thirty-2-character-string",
	}))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("6tq6VT3tOKcN1e6m6D5xtZfJi2EnNG7X"),
	}))

	err := app.Listen(":5000")
	if err != nil {
		fmt.Print(err)
	}
}

func CsrfFromHeader(name string) func(c *fiber.Ctx) (string, error) {

	return func(c *fiber.Ctx) (string, error) {
		return c.Get(name), nil
	}

}

func defaultErrorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "csrf token mismatch"})
}
