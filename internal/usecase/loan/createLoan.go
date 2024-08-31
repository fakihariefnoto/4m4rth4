package loan

import (
	"context"
	"fmt"
	"time"

	loanModel "billingapp/internal/model/loan"
)

func (l *loan) CreateLoan(loanReq LoanRequest) error {
	//set timezone,
	now := time.Now().In(location)
	// span time per week
	spanTime := time.Duration(24 * time.Hour * 7)
	amountTotal := (loanReq.Amount * AnnualRatePrecentage) + loanReq.Amount
	payablePerWeek := amountTotal / float64(NumberOfLoanWeeks)
	amountInterest := amountTotal - loanReq.Amount

	// total
	loan := loanModel.Loan{
		CustomerID:           loanReq.CustomerID,
		Name:                 loanReq.Name,
		Amount:               loanReq.Amount,
		AmountInterest:       amountInterest,
		AnnualRatePrecentage: AnnualRatePrecentage,
		Status:               loanModel.LoanStatusOngoing,
	}

	timeToPay := now
	var arrLoanDetails []loanModel.LoanDetails

	for i := 1; i <= NumberOfLoanWeeks; i++ {
		timeToPay = timeToPay.Add(spanTime)
		deadlinePayTime := timeToPay.Add(spanTime)

		if i == 1 {
			loan.StartDate = timeToPay
		}
		if i == NumberOfLoanWeeks {
			loan.EndDate = deadlinePayTime
		}

		weekFormat := fmt.Sprintf("Week-%d", i)

		arrLoanDetails = append(arrLoanDetails, loanModel.LoanDetails{
			Name:      weekFormat,
			Amount:    payablePerWeek,
			Status:    loanModel.LoanDetailStatusUnpaid,
			StartDate: timeToPay,
			EndDate:   deadlinePayTime,
		})
	}

	err := l.loanRepo.CreateLoanWithTx(context.Background(), loan, arrLoanDetails)
	if err != nil {
		return err
	}

	return nil

}
