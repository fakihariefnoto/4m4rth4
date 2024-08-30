package customer

import (
	// customerModel "billing/internal/model/customer/"
	customerRepository "github.com/fakihariefnoto/4m4rth4/module/repository/customer"
)

type (
	ICustomer interface {
	}

	customer struct {
		customerRepo customerRepository.ICustomer
	}
)

func New(customerRepo customerRepository.ICustomer) ICustomer {
	return &customer{
		customerRepo: customerRepo,
	}
}
