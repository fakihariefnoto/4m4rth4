package customer

type CreditStatusNum int

const (
	CreditStatusGood       CreditStatusNum = 0
	CreditStatusDelinquent CreditStatusNum = 1
)

type (
	Customer struct {
		ID           int64
		FullName     string
		Status       int
		CreditStatus CreditStatusNum
	}
)
