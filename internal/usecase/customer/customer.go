package customer

import (
	// customerModel "billing/internal/model/customer/"
	customerRepository "billingapp/internal/repository/customer"
	loanRepository "billingapp/internal/repository/loan"
)

type (
	ICustomer interface {
		GetCustomerData(customerID int64) (resp CustomerDetail, err error)
		IsDelinquent(customerID int64) (resp CustomerDelinquent, err error)
		CreateCustomer(name string) (resp CustomerDetail, err error)
	}

	customer struct {
		customerRepo customerRepository.ICustomer
		loanRepo     loanRepository.ILoan
	}
)

func New(customerRepo customerRepository.ICustomer, loanRepo loanRepository.ILoan) ICustomer {
	return &customer{
		customerRepo: customerRepo,
		loanRepo:     loanRepo,
	}
}
