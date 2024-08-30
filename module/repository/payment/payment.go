package payment

import (
	paymentModel "github.com/fakihariefnoto/4m4rth4/module/model/payment"

	"context"
	"database/sql"
)

type (
	IPayment interface {
		GetPaymentHistory(ctx context.Context, paymentID int64) (paymentModel.Payment, error)
		InsertPaymentHistory(ctx context.Context, arg paymentModel.Payment) (paymentID int64, err error)
		UpdatePaymentHistoryStatus(ctx context.Context, paymentID int64, status int) error
	}

	payment struct {
		db *sql.DB
	}
)

func New(paymentDB *sql.DB) IPayment {
	return &payment{
		db: paymentDB,
	}
}

func (q *payment) GetPaymentHistory(ctx context.Context, paymentID int64) (paymentModel.Payment, error) {
	row := q.db.QueryRowContext(ctx, queryGetPaymentHistoryByID, paymentID)
	var i paymentModel.Payment
	err := row.Scan(
		&i.ID,
		&i.Summary,
		&i.Details,
		&i.Amount,
		&i.Status,
	)
	return i, err
}

func (q *payment) InsertPaymentHistory(ctx context.Context, arg paymentModel.Payment) (paymentID int64, err error) {
	// Payment ID is autogenerated
	res, err := q.db.ExecContext(ctx, execInsertPaymentHistory,
		arg.Summary,
		arg.Details,
		arg.Amount,
		arg.Status,
	)
	if err != nil {
		return
	}
	paymentID, err = res.LastInsertId()
	return paymentID, err
}

func (q *payment) UpdatePaymentHistoryStatus(ctx context.Context, paymentID int64, status int) error {
	_, err := q.db.ExecContext(ctx, execUpdateCustomerCreditStatus, status, paymentID)
	return err
}
