package rest

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (p *presenter) GetCustomer(c *fiber.Ctx) error {

	customerIDstr := c.Params("customer_id")

	customerID, _ := strconv.ParseInt(customerIDstr, 10, 64)

	if customerID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "customer id not found",
			"data":  nil,
		})
	}

	resp, err := p.customerUse.GetCustomerData(customerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

func (p *presenter) IsDelinquent(c *fiber.Ctx) error {

	customerIDstr := c.Params("customer_id")

	customerID, _ := strconv.ParseInt(customerIDstr, 10, 64)

	if customerID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "customer id not found",
			"data":  nil,
		})
	}

	resp, err := p.customerUse.GetCustomerData(customerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
