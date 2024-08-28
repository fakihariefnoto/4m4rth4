package loan

import "time"

type LoanStatus int

const (
	LoanStatusInApproval LoanStatus = 0
	LoanStatusOngoing    LoanStatus = 1
	LoanStatusFinish     LoanStatus = 2
)

type LoanSDetailtatus int

const (
	LoanDetailStatusPending   LoanSDetailtatus = 0
	LoanDetailStatusPaid      LoanSDetailtatus = 1
	LoanDetailStatusCancelled LoanSDetailtatus = 2
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
		LoanID    int64
		Name      string
		Amount    float64
		Status    LoanSDetailtatus
		StartDate time.Time
		EndDate   time.Time
		PaymentID int64
	}
)
