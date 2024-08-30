package payment

import (
	//paymentModel "billing/internal/model/payment/"
	loanRepository "github.com/fakihariefnoto/4m4rth4/module/repository/loan"
	paymentRepository "github.com/fakihariefnoto/4m4rth4/module/repository/payment"
)

type (
	IPayment interface {
		MakePayment(customerID int64) error
	}

	payment struct {
		paymentRepo paymentRepository.IPayment
		loanRepo    loanRepository.ILoan
	}
)

func New(payRepo paymentRepository.IPayment, loanRep loanRepository.ILoan) IPayment {
	return &payment{
		paymentRepo: payRepo,
		loanRepo:    loanRep,
	}
}
