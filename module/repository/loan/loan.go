package loan

import (
	"fmt"
	"strings"

	loanModel "github.com/fakihariefnoto/4m4rth4/module/model/loan"

	"context"
	"database/sql"
)

type (
	ILoan interface {
		GetLoanByCustomerID(ctx context.Context, customerID int64, status loanModel.LoanStatus) (res []loanModel.Loan, err error)
		GetLoanByID(ctx context.Context, loanID int64) (loanModel.Loan, error)
		GetLoanDetailsByID(ctx context.Context, loanID int64) (loanModel.LoanDetail, error)
		CreateLoanWithTx(ctx context.Context, loan loanModel.Loan, arrLoan []loanModel.LoanDetails) error
		UpdateLoanStatus(ctx context.Context, loanID int64, status loanModel.PaymentStatus) error
		UpdateLoanDetailsStatus(ctx context.Context, loanDetailsID int64, status loanModel.PaymentStatus) error
	}

	loan struct {
		db *sql.DB
	}
)

func New(loanDB *sql.DB) ILoan {
	return &loan{
		db: loanDB,
	}
}

func (q *loan) GetLoanByCustomerID(ctx context.Context, customerID int64, status loanModel.LoanStatus) (res []loanModel.Loan, err error) {
	var rows *sql.Rows
	if status == loanModel.LoanStatusDefault {
		rows, err = q.db.QueryContext(ctx, queryGetLoanByCustomerIDAllStatus, customerID)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = q.db.QueryContext(ctx, queryGetLoanByCustomerID, customerID, status)
		if err != nil {
			return nil, err
		}
	}

	defer rows.Close()
	var items []loanModel.Loan
	for rows.Next() {
		var i loanModel.Loan
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.Name,
			&i.Amount,
			&i.AmountInterest,
			&i.AnnualRatePrecentage,
			&i.StartDate,
			&i.EndDate,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *loan) GetLoanByID(ctx context.Context, loanID int64) (loanModel.Loan, error) {
	row := q.db.QueryRowContext(ctx, queryGetLoanByID, loanID)
	var i loanModel.Loan
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.Name,
		&i.Amount,
		&i.AmountInterest,
		&i.AnnualRatePrecentage,
		&i.StartDate,
		&i.EndDate,
		&i.Status,
	)
	return i, err
}

func (q *loan) GetLoanDetailsByID(ctx context.Context, loanDetailsID int64) (loanModel.LoanDetail, error) {
	row := q.db.QueryRowContext(ctx, queryGetLoanDetails, loanDetailsID)
	var i loanModel.LoanDetails
	err := row.Scan(
		&i.ID,
		&i.LoanID,
		&i.Name,
		&i.Amount,
		&i.Status,
		&i.StartDate,
		&i.EndDate,
		&i.PaymentID,
	)
	return i, err
}

func (q *loan) GetLoanDetailsByLoanID(ctx context.Context, loanID int64) ([]loanModel.LoanDetails, error) {
	rows, err := q.db.QueryContext(ctx, queryGetLoanDetailsByLoanID, loanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []loanModel.LoanDetails
	for rows.Next() {
		var i loanModel.LoanDetails
		if err := rows.Scan(
			&i.ID,
			&i.LoanID,
			&i.Name,
			&i.Amount,
			&i.Status,
			&i.StartDate,
			&i.EndDate,
			&i.PaymentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *loan) CreateLoanWithTx(ctx context.Context, loan loanModel.Loan, arrLoan []loanModel.LoanDetails) error {
	tx, err := q.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	defer tx.Rollback()

	if err != nil {
		return err
	}

	id, err := q.insertLoan(ctx, tx, loan)
	if err != nil {
		return err
	}

	err = q.insertMultiLoanDetails(ctx, tx, id, arrLoan)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (q *loan) insertLoan(ctx context.Context, tx *sql.Tx, arg loanModel.Loan) (ID int64, err error) {
	if tx == nil {
		res, err := q.db.ExecContext(ctx, execInserLoan,
			arg.CustomerID,
			arg.Name,
			arg.Amount,
			arg.AmountInterest,
			arg.AnnualRatePrecentage,
			arg.StartDate,
			arg.EndDate,
			arg.Status,
		)

		if err != nil {
			return ID, err
		}

		return res.LastInsertId()
	}

	res, err := tx.ExecContext(ctx, execInserLoan,
		arg.ID,
		arg.CustomerID,
		arg.Name,
		arg.Amount,
		arg.AmountInterest,
		arg.AnnualRatePrecentage,
		arg.StartDate,
		arg.EndDate,
		arg.Status,
	)

	if err != nil {
		return ID, err
	}

	return res.LastInsertId()

}

func (q *loan) insertLoanDetails(ctx context.Context, arg loanModel.LoanDetails) error {
	_, err := q.db.ExecContext(ctx, execInsertLoanDetails,
		arg.LoanID,
		arg.Name,
		arg.Amount,
		arg.Status,
		arg.StartDate,
		arg.EndDate,
		arg.PaymentID,
	)
	return err
}

func (q *loan) insertMultiLoanDetails(ctx context.Context, tx *sql.Tx, loanID int64, arrLoan []loanModel.LoanDetails) error {
	fields := q.insertMultiLoanDetailsQueryBuilder(arrLoan, loanID)
	query := fmt.Sprintf(execInsertMultiLoanDetails, fields)

	if tx == nil {
		_, err := q.db.ExecContext(ctx, query)
		return err
	}

	_, err := tx.ExecContext(ctx, query)
	return err
}

func (q *loan) insertMultiLoanDetailsQueryBuilder(arrLoan []loanModel.LoanDetails, loanID int64) (query string) {
	var arrFields []string
	loanIDOverride := loanID

	for _, ld := range arrLoan {
		if loanID == 0 {
			loanIDOverride = ld.LoanID
		}
		tempField := fmt.Sprintf("(%d, %s, %.2f, %d, %v, %v, %d)", loanIDOverride, ld.Name, ld.Amount, ld.Status, ld.StartDate, ld.EndDate, ld.PaymentID)
		arrFields = append(arrFields, tempField)
	}
	return strings.Join(arrFields, ",")
}

func (q *loan) UpdateLoanStatus(ctx context.Context, loanID int64, status loanModel.PaymentStatus) error {
	_, err := q.db.ExecContext(ctx, execUpdateLoanStatus, status, loanID)
	return err
}

func (q *loan) UpdateLoanDetailsStatus(ctx context.Context, loanDetailsID int64, status loanModel.PaymentStatus) error {
	_, err := q.db.ExecContext(ctx, execUpdateLoanDetailsStatus,
		loanDetailsID,
		status,
	)
	return err
}
