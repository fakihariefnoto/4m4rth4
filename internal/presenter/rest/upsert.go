package rest

import (
	"github.com/gofiber/fiber/v2"
)

func (p *presenter) MakePayment(c *fiber.Ctx) error {

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  "",
	})
}

func (p *presenter) CreateLoan(c *fiber.Ctx) error {

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  "",
	})
}
