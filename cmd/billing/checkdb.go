package main

import (
	configPkg "billingapp/pkg/config"
	databasePkg "billingapp/service/database"
)

func checkDB(cfg configPkg.Config) {
	for _, db := range cfg.DB {
		dbNameSQL := "sql/" + db.Name + ".db"
		databasePkg.CreateDB(dbNameSQL)
		if databasePkg.ConnectionName(db.Name) == databasePkg.LoanDB {
			execCreateTable(dbNameSQL, tableLoan)
		}
		if databasePkg.ConnectionName(db.Name) == databasePkg.CustomerDB {
			execCreateTable(dbNameSQL, tableCustomer)
		}
		if databasePkg.ConnectionName(db.Name) == databasePkg.PaymentDB {
			execCreateTable(dbNameSQL, tablePayment)
		}
	}
}

func execCreateTable(dbName, syntax string) {
	db, err := databasePkg.Connect(dbName)
	chekErrMain("create table "+dbName+" err ", err)

	_, err = db.Exec(syntax)
	chekErrMain("exec create table "+dbName+" err ", err)
}

const (
	tableCustomer = `
CREATE TABLE IF NOT EXISTS customer (
    ID INTEGER PRIMARY KEY,
    full_name TEXT NOT NULL,
    status INTEGER,
    credit_status INTEGER
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_customer ON customer (ID);
`

	tableLoan = `
CREATE TABLE IF NOT EXISTS loan (
    ID INTEGER PRIMARY KEY,
    customer_id INTEGER,
    name TEXT,
    amount REAL,
    amount_interest REAL,
    annual_rate_precentage REAL,
    start_date DATE,
    end_date DATE,
    status INTEGER
);

CREATE INDEX IF NOT EXISTS idx_customer_loan ON loan (customer_id, status);
CREATE INDEX IF NOT EXISTS idx_customer_startdate_loan ON loan (customer_id, start_date, status);

CREATE TABLE IF NOT EXISTS loan_details (
	ID INTEGER PRIMARY KEY,
    loan_id INTEGER,
    name TEXT,
    amount REAL,
    status INTEGER,
    start_date DATE,
    end_date DATE,
    payment_id INTEGER,
    update_time DATE,
    FOREIGN KEY (loan_id) REFERENCES loan(ID)
);

CREATE INDEX IF NOT EXISTS idx_loanid_details_status ON loan_details (loan_id, status);
CREATE INDEX IF NOT EXISTS idx_loanid_details ON loan_details (loan_id, start_date);
`

	tablePayment = `
CREATE TABLE IF NOT EXISTS payment_history (
    payment_id INTEGER PRIMARY KEY,
    customer_id INTEGER,
    summary TEXT,
    details TEXT,
    amount REAL,
    status TEXT,
    update_time DATE
);
`
)
