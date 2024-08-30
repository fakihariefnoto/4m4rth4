package loan

import "database/sql"

type Loan struct {
	ID                   int64
	CustomerID           sql.NullInt64
	Name                 sql.NullString
	Ammount              sql.NullFloat64
	AmountInterest       sql.NullFloat64
	AnnualRatePrecentage sql.NullFloat64
	StartDate            sql.NullTime
	EndDate              sql.NullTime
	Status               sql.NullString
}

type LoanDetail struct {
	LoanID    sql.NullInt64
	WeekName  sql.NullString
	Amount    sql.NullFloat64
	Status    sql.NullString
	StartDate sql.NullTime
	EndDate   sql.NullTime
	PaymentID sql.NullInt64
}
