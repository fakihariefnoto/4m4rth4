package payment

import "database/sql"

type Customer struct {
	CustomerID   int64
	FullName     string
	Status       sql.NullString
	CreditStatus sql.NullString
}
