package rest

import (
	"log"

	"github.com/gofiber/fiber/v2"

	loanUsecase "billingapp/internal/usecase/loan"
)

func (p *presenter) CreateLoan(c *fiber.Ctx) error {

	var loan loanUsecase.LoanRequest

	if err := c.BodyParser(&loan); err != nil {
		log.Println("error when body parser ", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid request data",
			"data":  nil,
		})
	}

	err := p.loanUse.CreateLoan(loan)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"data":  nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  nil,
	})
}
