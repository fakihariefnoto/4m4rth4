package payment

import (
	//paymentModel "billing/internal/model/payment/"
	loanRepository "billingapp/internal/repository/loan"
	paymentRepository "billingapp/internal/repository/payment"
)

type (
	IPayment interface {
		MakePayment(paymentReq PaymentReq) error
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
