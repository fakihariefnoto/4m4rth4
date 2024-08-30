package loan

import "time"

type LoanStatus int

const (
	LoanStatusInApproval LoanStatus = 1
	LoanStatusOngoing    LoanStatus = 2
	LoanStatusFinish     LoanStatus = 3

	LoanStatusDefault LoanStatus = 0
)

type LoanDetailStatus int

const (
	LoanDetailStatusUnpaid    LoanDetailStatus = 1
	LoanDetailStatusPaid      LoanDetailStatus = 2
	LoanDetailStatusCancelled LoanDetailStatus = 3

	LoanDetailStatusDefault LoanDetailStatus = 0
)

type (
	Loan struct {
		ID                   int64
		CustomerID           int64
		Name                 string
		Amount               float64
		AmountInterest       float64
		AnnualRatePrecentage float64
		StartDate            time.Time
		EndDate              time.Time
		Status               LoanStatus
	}

	LoanDetails struct {
		ID        int64
		LoanID    int64
		Name      string
		Amount    float64
		Status    LoanDetailStatus
		StartDate time.Time
		EndDate   time.Time
		PaymentID int64
	}
)
