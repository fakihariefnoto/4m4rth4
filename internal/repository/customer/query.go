package customer

const (
	queryGetCustomerByID = `
		SELECT
			customer_id, full_name, status, credit_status 
		FROM
			customer
		WHERE
			customer_id = ?
	`

	execCustomer = `
		INSERT INTO 
			customer (customer_id, full_name, status, credit_status)
		VALUES
			(?, ?, ?, ?)
	`

	execUpdateCustomerCreditStatus = `
		UPDATE 
			customer
		SET
			credit_status = ?
		WHERE
			customer_id = ?
	`
)
