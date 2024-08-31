package customer

type CustomerStatus int

const (
	CustomerStatusInactive CustomerStatus = 0
	CustomerStatusActive   CustomerStatus = 1
	CustomerStatusDef      CustomerStatus = 1
)

type CreditStatusNum int

const (
	CreditStatusGood       CreditStatusNum = 1
	CreditStatusDelinquent CreditStatusNum = 2
	CreditStatusDef        CreditStatusNum = 0
)

type (
	Customer struct {
		ID           int64
		FullName     string
		Status       CustomerStatus
		CreditStatus CreditStatusNum
	}
)
