package route

import (
	"github.com/AdrianYuu/easy-laundry-be/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	Fiber          *fiber.App
	UserController *http.UserController
}

func (c *RouteConfig) Setup() {
	c.Fiber.Get("/api/ping", c.UserController.Ping)
}
