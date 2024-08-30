package payment

type (
	PaymentReq struct {
		CustomerID    int64   `json:"customer_id"`
		LoanDetailsID int64   `json:"loan_detail_id"`
		Amount        float64 `json:"amount"`
	}
)
