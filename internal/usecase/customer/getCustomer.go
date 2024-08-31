package customer

import (
	customerModel "billingapp/internal/model/customer"
	"context"
	"errors"
)

func (c *customer) GetCustomerData(customerID int64) (resp CustomerDetail, err error) {
	if customerID == 0 {
		return resp, errors.New("Invalid customer ID")
	}

	data, err := c.customerRepo.GetCustomer(context.Background(), customerID)
	if err != nil {
		return resp, err
	}

	resp = CustomerDetail{
		ID:           data.ID,
		FullName:     data.FullName,
		Status:       CustomerStatusToString(data.Status),
		CreditStatus: CreditStatusToString(data.CreditStatus),
	}

	return
}

func (c *customer) IsDelinquent(customerID int64) (resp CustomerDelinquent, err error) {
	if customerID != 0 {
		return resp, errors.New("Invalid customer ID")
	}

	data, err := c.customerRepo.GetCustomer(context.Background(), customerID)
	if err != nil {
		return resp, err
	}

	isDelinquent := false
	if data.CreditStatus == customerModel.CreditStatusDelinquent {
		isDelinquent = true
	}

	resp = CustomerDelinquent{
		IsDelinquent: isDelinquent,
	}

	return
}
