package middleware

import (
	"github.com/donhuvy/go-rest-api-with-fiber/config"
	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
)

// AuthReq middleware
func AuthReq() func(*fiber.Ctx) {
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}
	err := basicauth.New(cfg)
	return err
}
