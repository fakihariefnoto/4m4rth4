package customer

import (
	customerModel "billingapp/internal/model/customer"
	loanModel "billingapp/internal/model/loan"
	loanUsecase "billingapp/internal/usecase/loan"

	"context"
	"errors"
	"time"
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
	if customerID == 0 {
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

	dataCust, err := c.getLoanListByCustomerID(customerID)
	if err != nil {
		return resp, err
	}

	for _, value := range dataCust {
		if value.OverDueCounter >= 2 {
			isDelinquent = true
		}
	}

	resp = CustomerDelinquent{
		IsDelinquent: isDelinquent,
	}

	return
}

func (c *customer) getLoanListByCustomerID(ID int64) (resp []loanUsecase.LoanData, err error) {
	if ID == 0 {
		return nil, errors.New("Customer ID not found")
	}

	var getFirstStatus loanModel.LoanStatus

	data, err := c.loanRepo.GetLoanByCustomerID(context.Background(), ID, getFirstStatus)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	for _, loanData := range data {
		startDate := loanData.StartDate.Format("2006-01-02 15:04:05")
		endDate := loanData.EndDate.Format("2006-01-02 15:04:05")

		var overdueCounter int
		dataDetails, err := c.loanRepo.GetLoanDetailsByLoanID(context.Background(), loanData.ID)
		if err != nil {
			return nil, err
		}
		for _, loanData := range dataDetails {
			status := loanUsecase.LoanDetailsStatusToString(loanData.Status)
			if status != loanUsecase.LoanDetailStatusPaidString {
				if loanData.EndDate.Before(now) {
					status = loanUsecase.LoanDetailStatusOverDueString
				}
				if loanData.StartDate.Before(now) && loanData.EndDate.After(now) {
					status = loanUsecase.LoanDetailStatusUnpaidString
				}
				if loanData.StartDate.After(now) {
					status = loanUsecase.LoanDetailStatusNotDueString
				}
			}

			if status == loanUsecase.LoanDetailStatusOverDueString {
				overdueCounter++
			}
		}

		resp = append(resp, loanUsecase.LoanData{
			ID:                   loanData.ID,
			CustomerID:           loanData.CustomerID,
			Name:                 loanData.Name,
			Amount:               loanData.Amount,
			AmountInterest:       loanData.AmountInterest,
			AnnualRatePrecentage: loanData.AnnualRatePrecentage,
			StartDate:            startDate,
			EndDate:              endDate,
			Status:               loanUsecase.LoanStatusToString(loanData.Status),
			OverDueCounter:       overdueCounter,
		})
	}

	return resp, nil
}
