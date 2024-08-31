# Billing
created for 4m4rth4 assignment test example-1 (billing engine)

## Functionality

### All Method That mention in the docs for example-1 (billing engine)
- GetOutstanding : This returns the current outstanding on a loan, 0 if no outstanding(or closed)
```
GET api/v1/GetOutstanding/:customer_id
```
- IsDelinquent : If there are more than 2 weeks of Non payment of the loan amount
```
GET api/v1/IsDelinquent/:customer_id
```
- MakePayment: Make a payment of certain amount on the loan
```
POST api/v1/MakePayment
```

### Additional Method
- CreateLoan : Creating Loan by CustomerID
```
POST api/v1/CreateLoan
```
- ListLoan : List of Loan by CustomerID (On Going or Finished)
```
GET api/v1/ListLoan/:customer_id
```
- GetLoanDetails : Get Loan Details by Loan ID
```
GET api/v1/GetLoanDetails/:loan_id
```
- GetCustomer : Get Customer Details
```
GET api/v1/GetCustomer/:customer_id
```
- CreateCustomer : Create new customer
```
POST api/v1/CreateCustomer
```

### Cronjob to update status of customer

## How To Build/Run

### Golang
#### Prerequiested

- Go 1.21 above already installed 

#### Build & Run

Run this command
`make build & make run`

### Docker



## How To Test

### Existing Customer Journey

- CreateCustomer

- CreateLoan

- ListLoan

- GetOutstanding

- IsDelinquent / GetCustomer

- GetLoanDetails

- MakePayment

### New Customer Journey

- CreateLoan

- ListLoan

- GetOutstanding

- IsDelinquent / GetCustomer

- GetLoanDetails

- MakePayment

## Architecture

### Code Project Structure

```
├── bin
├── cmd
│   └── billing
├── config
├── handler
│   ├── cron
│   └── http
├── internal
│   ├── model
│   │   ├── customer
│   │   ├── loan
│   │   └── payment
│   ├── presenter
│   │   ├── cron
│   │   └── rest
│   ├── repository
│   │   ├── customer
│   │   ├── loan
│   │   └── payment
│   └── usecase
│       ├── customer
│       ├── loan
│       └── payment
├── log
├── pkg
│   └── config
├── service
│   ├── cache
│   └── database
└── sql
```

### Model

#### Loan

```
CREATE TABLE IF NOT EXISTS loan (
    ID INTEGER PRIMARY KEY,
    customer_id INTEGER,
    name TEXT,
    amount REAL,
    amount_interest REAL,
    annual_rate_precentage REAL,
    start_date DATE,
    end_date DATE,
    status INTEGER,
    update_time DATE
);

CREATE INDEX IF NOT EXISTS idx_customer_loan ON loan (customer_id, status);
CREATE INDEX IF NOT EXISTS idx_customer_startdate_loan ON loan (customer_id, start_date, status);
```

#### LoanDetail

```
CREATE TABLE IF NOT EXISTS loan_details (
	ID INTEGER PRIMARY KEY,
    loan_id INTEGER,
    name TEXT,
    amount REAL,
    status INTEGER,
    start_date DATE,
    end_date DATE,
    payment_id INTEGER,
    update_time DATE
);

CREATE INDEX IF NOT EXISTS idx_loanid_details_status ON loan_details (loan_id, status);
CREATE INDEX IF NOT EXISTS idx_loanid_details ON loan_details (loan_id, start_date);
```

#### Payment
```
CREATE TABLE IF NOT EXISTS payment_history (
    payment_id INTEGER PRIMARY KEY,
    customer_id INTEGER,
    summary TEXT,
    details TEXT,
    amount REAL,
    status INTEGER,
    create_time DATE
);
```

#### Customer
```
CREATE TABLE IF NOT EXISTS customer (
    ID INTEGER PRIMARY KEY,
    full_name TEXT NOT NULL,
    status INTEGER,
    credit_status INTEGER
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_customer ON customer (ID);
```

