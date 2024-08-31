package loan

const (
	queryGetLoanByCustomerID = `
		SELECT
			ID, 
			customer_id, 
			name, 
			amount, 
			amount_interest, 
			annual_rate_precentage, 
			start_date, 
			end_date, 
			status 
		FROM
			loan
		WHERE
			customer_id = ? AND status = ? ORDER BY start_date ASC
	`

	queryGetLoanByCustomerIDAllStatus = `
		SELECT
			ID, 
			customer_id, 
			name, 
			amount, 
			amount_interest, 
			annual_rate_precentage, 
			start_date, 
			end_date, 
			status 
		FROM
			loan
		WHERE
			customer_id = ? ORDER BY start_date ASC
	`

	queryGetLoanByID = `
		SELECT
			ID, 
			customer_id, 
			name, 
			amount, 
			amount_interest, 
			annual_rate_precentage, 
			start_date, 
			end_date, 
			status 
		FROM
			loan
		WHERE
			id = ?
	`

	queryGetLoanDetails = `
		SELECT
			ID,
			loan_id,
			name,
			amount,
			status,
			start_date,
			end_date,
			payment_id
		FROM
			loan_details
		WHERE
			id = ?`

	queryGetLoanDetailsByLoanID = `
		SELECT
			ID,
			loan_id,
			name,
			amount,
			status,
			start_date,
			end_date,
			payment_id
		FROM
			loan_details
		WHERE
			loan_id = ?`

	execInserLoan = `
		INSERT INTO 
			loan (customer_id, name, amount, amount_interest, annual_rate_precentage, start_date, end_date, status)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id`

	execInsertLoanDetails = `
		INSERT INTO 
			loan_details (loan_id, name, amount, status, start_date, end_date, payment_id)
		VALUES
			(?, ?, ?, ?, ?, ?, ?)`

	execInsertMultiLoanDetails = `
		INSERT INTO 
			loan_details (loan_id, name, amount, status, start_date, end_date, payment_id)
		VALUES
			%s`

	execUpdateLoanDetailsStatus = `
		UPDATE
			loan_details
		SET
			status = ?,
			payment_id = ?
		WHERE
			loan_id = ?`

	execUpdateLoanStatus = `
		UPDATE
			loan
		SET
			status ?
		WHERE
			ID = ?;`
)
