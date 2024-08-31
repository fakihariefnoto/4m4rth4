package customer

import (
	customerModel "billingapp/internal/model/customer"
	"context"
	"errors"
)

func (c *customer) CreateCustomer(name string) (resp CustomerDetail, err error) {
	if name == "" {
		return resp, errors.New("Invalid customer name")
	}

	customerID, err := c.customerRepo.InsertCustomer(context.Background(), customerModel.Customer{
		FullName:     name,
		Status:       customerModel.CustomerStatusActive,
		CreditStatus: customerModel.CreditStatusGood,
	})
	if err != nil {
		return resp, err
	}

	resp = CustomerDetail{
		ID: customerID,
	}

	return
}
