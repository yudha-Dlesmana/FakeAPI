package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yudha-Dlesmana/fakeAPI/types"
)

// welcomehandler godoc
// @Tags					Welcome
// @Summary				Welcome to FakeAPI
// @Description		Basic information about the FakeAPI service, including status, version
// @Success 			200 {object} types.DefaultResponse
// @Router				/v1 [get]
func WelcomeHandler(c *fiber.Ctx) error {
	response := types.DefaultResponse{
		Message: "Welcome to FakeAPI",
		Info:    "This is a fake REST API for prototyping, frontend development, or testing purposes only.",
		Status:  "development",
		Version: "v1",
	}
	return c.JSON(response)
}
