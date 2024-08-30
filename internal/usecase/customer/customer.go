package customer

import (
	// customerModel "billing/internal/model/customer/"
	customerRepository "billingapp/internal/repository/customer"
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
