package customer

import (
	"context"
	"database/sql"

	customerModel "billingapp/internal/model/customer"
)

type (
	ICustomer interface {
		GetCustomer(ctx context.Context, ID int64) (customerModel.Customer, error)
		InsertCustomer(ctx context.Context, arg customerModel.Customer) (int64, error)
		UpdateCustomerCreditStatus(ctx context.Context, ID int64, creditStatus string) error
	}

	customer struct {
		db *sql.DB
	}
)

func New(customerDB *sql.DB) ICustomer {
	return &customer{
		db: customerDB,
	}
}
func (q *customer) GetCustomer(ctx context.Context, ID int64) (customerModel.Customer, error) {
	row := q.db.QueryRowContext(ctx, queryGetCustomerByID, ID)
	var i customerModel.Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Status,
		&i.CreditStatus,
	)
	return i, err
}

func (q *customer) InsertCustomer(ctx context.Context, arg customerModel.Customer) (int64, error) {
	// ID will be auto increment
	tx, err := q.db.Begin()
	if err != nil {
		return 0, err
	}

	res := tx.QueryRow(execCustomer,
		arg.FullName,
		arg.Status,
		arg.CreditStatus,
	)

	var id int64

	err = res.Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return id, err
	}
	return id, err
}

func (q *customer) UpdateCustomerCreditStatus(ctx context.Context, ID int64, creditStatus string) error {
	_, err := q.db.ExecContext(ctx, execUpdateCustomerCreditStatus, creditStatus, ID)
	return err
}
