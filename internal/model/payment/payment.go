package payment

type PaymentStatus int

const (
	PaymentStatusReceived PaymentStatus = 0
	PaymentStatusRejected PaymentStatus = 1
	PaymentStatusExpired  PaymentStatus = 1
)

type (
	Payment struct {
		ID         int64
		CustomerID int64
		Details    string
		Summary    string
		Amount     float64
		Status     PaymentStatus
	}
)
