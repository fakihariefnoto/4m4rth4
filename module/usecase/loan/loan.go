package loan

import (
	loanModel "github.com/fakihariefnoto/4m4rth4/module/model/loan"
	loanRepository "github.com/fakihariefnoto/4m4rth4/module/repository/loan"

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
		loanRepo loanRepository.Iloan
	}
)

func New(loanRepo loanRepository.ILoan) ILoan {
	location, _ = time.LoadLocation("Asia/Jakarta")
	return &loan{
		loanRepo: loanRepo,
	}
}
