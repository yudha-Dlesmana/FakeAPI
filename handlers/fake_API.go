package handlers

import (
	"github.com/yudha-Dlesmana/fakeAPI/services"
	"github.com/yudha-Dlesmana/fakeAPI/types"

	"github.com/gofiber/fiber/v2"
)

// FakeAPIHandler godoc
// @Summary 			Generate Fake API Data
// @Description		Generate fake data base on the provided type and count parameters.
// @Tags					FakeAPI
// @Accept				json
// @Produce				json
// @Param					type header string true "Faker type mapping (e.g. {\'user\':\'username\', ...})"
// @Param					count query int false "Number of fake data to generate (default 1)"
// @Success				200 {object} types.Response
// @Failure				400 {object} types.Response
// @Router				/v1/fakeAPI [get]
func FakeApiHandler(c *fiber.Ctx) error {
	reqType := c.Get("type")
	reqCount := c.QueryInt("count", 1)

	fakeAPI, err := services.GenerateFakeData(reqType, reqCount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			types.Response{
				Status:  400,
				Message: err.Error(),
				Data:    "",
			})
	}

	return c.JSON(types.Response{
		Status:  200,
		Message: "success",
		Data:    fakeAPI,
	})
}
