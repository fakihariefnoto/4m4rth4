package loan

import (
	"time"

	loanModel "github.com/fakihariefnoto/4m4rth4/module/model/loan"
)

const (
	NumberOfLoanWeeks    int     = 50
	AnnualRatePrecentage float64 = 10
)

var location *time.Location

type (
	LoanData struct {
		ID                   int64             `json:"id"`
		CustomerID           int64             `json:"customer_id"`
		Name                 string            `json:"name"`
		Amount               float64           `json:"amount"`
		AmountInterest       float64           `json:"amount_interest"`
		AnnualRatePrecentage float64           `json:"rate"`
		StartDate            string            `json:"start_date"`
		EndDate              string            `json:"end_date"`
		Status               LoanStatusString  `json:"status"`
		Details              []LoanDataDetails `json:"details"`
		TotalPaid            float64           `json:"total_paid"`
		TotalBorrowed        float64           `json:"total_borrowed"`
	}

	LoanDataDetails struct {
		ID        int64                  `json:"id"`
		LoanID    int64                  `json:"loan_id"`
		Name      string                 `json:"name"`
		Amount    float64                `json:"amount"`
		StartDate string                 `json:"start_date"`
		EndDate   string                 `json:"end_date"`
		Status    LoanDetailStatusString `json:"status"`
		PaymentID int64                  `json:"payment_id"`
	}

	LoanRequest struct {
		CustomerID int64   `json:"customer_id"`
		Name       string  `json:"name"`
		Amount     float64 `json:"amount"`
	}
)

type LoanStatusString string
type LoanDetailStatusString string

const (
	LoanStatusInApprovalString LoanStatusString = "InApproval"
	LoanStatusOngoingString    LoanStatusString = "OnGoing"
	LoanStatusFinishString     LoanStatusString = "Finished"

	LoanStatusEmptyString LoanStatusString = ""

	LoanDetailStatusUnpaidString    LoanDetailStatusString = "Unpaid"
	LoanDetailStatusPaidString      LoanDetailStatusString = "Paid"
	LoanDetailStatusCancelledString LoanDetailStatusString = "Cancelled"
	LoanDetailStatusNotDueString    LoanDetailStatusString = "NotDue"
	LoanDetailStatusOverDueString   LoanDetailStatusString = "OverDue"

	LoanDetailStatusEmptyString LoanDetailStatusString = ""
)

func LoanStatusToString(status loanModel.LoanStatus) LoanStatusString {
	switch status {
	case loanModel.LoanStatusInApproval:
		return LoanStatusInApprovalString
	case loanModel.LoanStatusOngoing:
		return LoanStatusOngoingString
	case loanModel.LoanStatusFinish:
		return LoanStatusFinishString
	default:
		return LoanStatusEmptyString
	}
}

func LoanStatusFromString(status LoanStatusString) loanModel.LoanStatus {
	switch status {
	case LoanStatusInApprovalString:
		return loanModel.LoanStatusInApproval
	case LoanStatusOngoingString:
		return loanModel.LoanStatusOngoing
	case LoanStatusFinishString:
		return loanModel.LoanStatusFinish
	default:
		return loanModel.LoanStatus
	}
}

func LoanDetailsStatusToString(status loanModel.LoanDetailStatus) LoanDetailStatusString {
	switch status {
	case loanModel.LoanDetailStatusUnpaid:
		return LoanDetailStatusUnpaidString
	case loanModel.LoanDetailStatusPaid:
		return LoanDetailStatusPaidString
	case loanModel.LoanDetailStatusCancelled:
		return LoanDetailStatusCancelledString
	default:
		return LoanDetailStatusEmptyString
	}
}

func LoanDetailsStatusFromString(status LoanDetailStatusString) loanModel.LoanDetailStatus {
	switch status {
	case LoanDetailStatusUnpaidString:
		return loanModel.LoanDetailStatusUnpaid
	case LoanDetailStatusPaidString:
		return loanModel.LoanDetailStatusPaid
	case LoanDetailStatusCancelledString:
		return loanModel.LoanDetailStatusCancelled
	default:
		return loanModel.LoanDetailStatus
	}
}
