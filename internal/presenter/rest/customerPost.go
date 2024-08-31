package rest

import (
	"log"

	customerUsecase "billingapp/internal/usecase/customer"

	"github.com/gofiber/fiber/v2"
)

func (p *presenter) CreateCustomer(c *fiber.Ctx) error {

	var cust customerUsecase.CostumerRequest

	if err := c.BodyParser(&cust); err != nil {
		log.Println("error when body parser ", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid request data",
			"data":  nil,
		})
	}

	if cust.FullName == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid request customer name",
			"data":  nil,
		})
	}

	resp, err := p.customerUse.CreateCustomer(*cust.FullName)
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
		"data":  resp,
	})
}
