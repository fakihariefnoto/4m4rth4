package rest

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (p *presenter) GetOutstanding(c *fiber.Ctx) error {
	loanIDstr := c.Params("loan_id")

	loanID, _ := strconv.ParseInt(loanIDstr, 10, 64)

	if loanID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "loan id not valid",
			"data":  nil,
		})
	}

	resp, err := p.loanUse.GetLoanByID(loanID)
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

func (p *presenter) GetLoanDetails(c *fiber.Ctx) error {

	loanIDstr := c.Params("loan_id")

	loanID, _ := strconv.ParseInt(loanIDstr, 10, 64)

	if loanID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "loan id not valid",
			"data":  nil,
		})
	}

	resp, err := p.loanUse.GetLoanByID(loanID)
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

func (p *presenter) ListLoan(c *fiber.Ctx) error {

	customerIDstr := c.Params("customer_id")

	customerID, _ := strconv.ParseInt(customerIDstr, 10, 64)

	if customerID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "loan id not valid",
			"data":  nil,
		})
	}

	resp, err := p.loanUse.GetLoanListByCustomerID(customerID)
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
