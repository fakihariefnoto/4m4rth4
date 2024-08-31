package loan

import (
	"context"
	"errors"
	"time"

	loanModel "billingapp/internal/model/loan"
)

func (l *loan) GetLoanListByCustomerID(ID int64, status ...loanModel.LoanStatus) (resp []LoanData, err error) {
	if ID == 0 {
		return nil, errors.New("Customer ID not found")
	}

	var getFirstStatus loanModel.LoanStatus
	if len(status) > 0 {
		getFirstStatus = status[0]
	}

	data, err := l.loanRepo.GetLoanByCustomerID(context.Background(), ID, getFirstStatus)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	for _, loanData := range data {
		startDate := loanData.StartDate.Format("2006-01-02 15:04:05")
		endDate := loanData.EndDate.Format("2006-01-02 15:04:05")

		var overdueCounter int
		dataDetails, err := l.loanRepo.GetLoanDetailsByLoanID(context.Background(), loanData.ID)
		if err != nil {
			return nil, err
		}
		for _, loanData := range dataDetails {
			status := LoanDetailsStatusToString(loanData.Status)
			if status != LoanDetailStatusPaidString {
				if loanData.EndDate.Before(now) {
					status = LoanDetailStatusOverDueString
				}
				if loanData.StartDate.Before(now) && loanData.EndDate.After(now) {
					status = LoanDetailStatusUnpaidString
				}
				if loanData.StartDate.After(now) {
					status = LoanDetailStatusNotDueString
				}
			}

			if status == LoanDetailStatusOverDueString {
				overdueCounter++
			}
		}

		resp = append(resp, LoanData{
			ID:                   loanData.ID,
			CustomerID:           loanData.CustomerID,
			Name:                 loanData.Name,
			Amount:               loanData.Amount,
			AmountInterest:       loanData.AmountInterest,
			AnnualRatePrecentage: loanData.AnnualRatePrecentage,
			StartDate:            startDate,
			EndDate:              endDate,
			Status:               LoanStatusToString(loanData.Status),
			OverDueCounter:       overdueCounter,
		})
	}

	return resp, nil
}

func (l *loan) GetLoanByID(ID int64) (resp LoanData, err error) {
	if ID == 0 {
		return resp, errors.New("Loan ID not found")
	}

	data, err := l.loanRepo.GetLoanByID(context.Background(), ID)
	startDate := data.StartDate.Format("2006-01-02 15:04:05")
	endDate := data.EndDate.Format("2006-01-02 15:04:05")

	resp = LoanData{
		ID:                   data.ID,
		Name:                 data.Name,
		CustomerID:           data.CustomerID,
		Amount:               data.Amount,
		AmountInterest:       data.AmountInterest,
		AnnualRatePrecentage: data.AnnualRatePrecentage,
		StartDate:            startDate,
		EndDate:              endDate,
		Status:               LoanStatusToString(data.Status),
		TotalBorrowed:        data.Amount + data.AmountInterest,
	}

	now := time.Now()

	dataDetails, err := l.loanRepo.GetLoanDetailsByLoanID(context.Background(), resp.ID)
	for _, loanData := range dataDetails {
		startDate := loanData.StartDate.Format("2006-01-02 15:04:05")
		endDate := loanData.EndDate.Format("2006-01-02 15:04:05")

		status := LoanDetailsStatusToString(loanData.Status)
		if status != LoanDetailStatusPaidString {
			if loanData.EndDate.Before(now) {
				status = LoanDetailStatusOverDueString
			}
			if loanData.StartDate.Before(now) && loanData.EndDate.After(now) {
				status = LoanDetailStatusUnpaidString
			}
			if loanData.StartDate.After(now) {
				status = LoanDetailStatusNotDueString
			}
		}

		if status == LoanDetailStatusPaidString {
			resp.TotalPaid += loanData.Amount
		}

		resp.Details = append(resp.Details, LoanDataDetails{
			ID:        loanData.ID,
			LoanID:    loanData.LoanID,
			Name:      loanData.Name,
			Amount:    loanData.Amount,
			StartDate: startDate,
			EndDate:   endDate,
			Status:    status,
		})
	}

	resp.OutStanding = resp.TotalBorrowed - resp.TotalPaid

	return
}
