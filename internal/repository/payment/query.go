package payment

const (
	queryGetPaymentHistoryByID = `
		SELECT
			payment_id, summary, details, amount, status 
		FROM
			payment_history
		WHERE
			payment_id = ?
	`

	execInsertPaymentHistory = `
		INSERT INTO 
			payment_history (summary, details, amount, status)
		VALUES
			(?, ?, ?, ?)
		RETURNING
			payment_id
	`

	execUpdateCustomerCreditStatus = `
		UPDATE
			payment_history
		SET
			status = ?
		WHERE
			payment_id = ?
	`
)
