package customer

const (
	queryGetCustomerByID = `
		SELECT
			ID, full_name, status, credit_status 
		FROM
			customer
		WHERE
			ID = ?
	`

	execCustomer = `
		INSERT INTO 
			customer (full_name, status, credit_status)
		VALUES
			(?, ?, ?)
		RETURNING ID
	`

	execUpdateCustomerCreditStatus = `
		UPDATE 
			customer
		SET
			credit_status = ?
		WHERE
			ID = ?
	`
)
