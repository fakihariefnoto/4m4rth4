package rest

import (
	"log"

	"github.com/gofiber/fiber/v2"

	paymentUsecase "billingapp/internal/usecase/payment"
)

func (p *presenter) MakePayment(c *fiber.Ctx) error {

	var pay paymentUsecase.PaymentReq

	if err := c.BodyParser(&pay); err != nil {
		log.Println("error when body parser ", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid request data",
			"data":  nil,
		})
	}

	err := p.paymentUse.MakePayment(pay)
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
