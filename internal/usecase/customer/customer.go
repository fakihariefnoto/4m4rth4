package customer

import (
	// customerModel "billing/internal/model/customer/"
	customerRepository "billingapp/internal/repository/customer"
)

type (
	ICustomer interface {
		GetCustomerData(customerID int64) (resp CustomerDetail, err error)
		IsDelinquent(customerID int64) (resp CustomerDelinquent, err error)
		CreateCustomer(name string) (resp CustomerDetail, err error)
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
