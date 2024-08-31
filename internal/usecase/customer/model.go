package customer

import (
	customerModel "billingapp/internal/model/customer"
)

type (
	CustomerDetail struct {
		ID           int64             `json:"id,omitempty"`
		FullName     string            `json:"full_name,omitempty"`
		Status       CustomerStatusStr `json:"status,omitempty"`
		CreditStatus CreditStatusStr   `json:"credit_status,omitempty"`
	}

	CustomerDelinquent struct {
		IsDelinquent bool `json:"is_delinquent"`
	}

	CostumerRequest struct {
		ID           *int64             `json:"id"`
		FullName     *string            `json:"full_name"`
		Status       *CustomerStatusStr `json:"status"`
		CreditStatus *CreditStatusStr   `json:"credit_status"`
	}
)

type CustomerStatusStr string

const (
	CustomerStatusActiveStr   CustomerStatusStr = "Active"
	CustomerStatusInactiveStr CustomerStatusStr = "Inactive"
	CustomerStatusDefStr      CustomerStatusStr = ""
)

type CreditStatusStr string

const (
	CreditStatusGoodStr       CreditStatusStr = "Good"
	CreditStatusDelinquentStr CreditStatusStr = "Delinquent"
	CreditStatusDefStr        CreditStatusStr = ""
)

func CustomerStatusToString(status customerModel.CustomerStatus) CustomerStatusStr {
	switch status {
	case customerModel.CustomerStatusActive:
		return CustomerStatusActiveStr
	case customerModel.CustomerStatusInactive:
		return CustomerStatusInactiveStr
	default:
		return CustomerStatusDefStr
	}
}

func CustomerStatusFromString(status CustomerStatusStr) customerModel.CustomerStatus {
	switch status {
	case CustomerStatusActiveStr:
		return customerModel.CustomerStatusActive
	case CustomerStatusInactiveStr:
		return customerModel.CustomerStatusInactive
	default:
		return customerModel.CustomerStatusDef
	}
}

func CreditStatusToString(status customerModel.CreditStatusNum) CreditStatusStr {
	switch status {
	case customerModel.CreditStatusGood:
		return CreditStatusGoodStr
	case customerModel.CreditStatusDelinquent:
		return CreditStatusDelinquentStr
	default:
		return CreditStatusDefStr
	}
}

func CreditStatusFromString(status CreditStatusStr) customerModel.CreditStatusNum {
	switch status {
	case CreditStatusGoodStr:
		return customerModel.CreditStatusGood
	case CreditStatusDelinquentStr:
		return customerModel.CreditStatusDelinquent
	default:
		return customerModel.CreditStatusDef
	}
}
