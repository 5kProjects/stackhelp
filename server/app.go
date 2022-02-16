package server

import (
	"github.com/gofiber/fiber/v2"
	"helpstack/config"
)

type MainApp struct {
	Engine  *fiber.App
	Configs config.AppConfig
}
