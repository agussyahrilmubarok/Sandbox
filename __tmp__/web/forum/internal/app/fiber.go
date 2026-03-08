package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(cfg *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName: cfg.GetString("APPLICATION.NAME"),
		ErrorHandler: newErrorHandler(),
		Prefork: cfg.GetBool("WEB.PREFORK"),
	})

	return app
}

func newErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
