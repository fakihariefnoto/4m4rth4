package loan

import (
	loanModel "billingapp/internal/model/loan"
	loanRepository "billingapp/internal/repository/loan"

	"time"
)

type (
	ILoan interface {
		GetLoanListByCustomerID(ID int64, status loanModel.LoanStatus) (resp []LoanData, err error)
		CreateLoan(loanReq LoanRequest) error
		GetLoanByID(ID int64) (resp LoanData, err error)
		UpdateLoanStatus(loanReq LoanRequest) error
		UpdateLoanDetailStatus(loanReq LoanRequest) error
	}

	loan struct {
		loanRepo loanRepository.ILoan
	}
)

func New(loanRepo loanRepository.ILoan) ILoan {
	location, _ = time.LoadLocation("Asia/Jakarta")
	return &loan{
		loanRepo: loanRepo,
	}
}
