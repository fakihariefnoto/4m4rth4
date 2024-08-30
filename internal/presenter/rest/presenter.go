package rest

import (
	"github.com/gofiber/fiber/v2"

	customerUsecase "billingapp/internal/usecase/customer"
	loanUsecase "billingapp/internal/usecase/loan"
	paymentUsecase "billingapp/internal/usecase/payment"
)

type (
	presenter struct {
		loanUse     loanUsecase.ILoan
		paymentUse  paymentUsecase.IPayment
		customerUse customerUsecase.ICustomer
	}

	IPresenter interface {
		GetOutstanding(c *fiber.Ctx) error
		IsDelinquent(c *fiber.Ctx) error
		MakePayment(c *fiber.Ctx) error
		CreateLoan(c *fiber.Ctx) error
		ListLoan(c *fiber.Ctx) error
		GetLoanDetails(c *fiber.Ctx) error
		GetCustomer(c *fiber.Ctx) error
	}
)

func New(loanUse loanUsecase.ILoan, paymentUse paymentUsecase.IPayment, customerUse customerUsecase.ICustomer) IPresenter {
	return &presenter{
		loanUse:     loanUse,
		paymentUse:  paymentUse,
		customerUse: customerUse,
	}
}
