# Billing
created for 4m4rth4 assignment test example-1 (billing engine)

## Functionality

### All Method That mention in the docs for example-1 (billing engine)
- GetOutstanding : This returns the current outstanding on a loan, 0 if no outstanding(or closed)
- IsDelinquent : If there are more than 2 weeks of Non payment of the loan amount
- MakePayment: Make a payment of certain amount on the loan

Additional Method
- CreateLoan : Creating Loan by CustomerID
- ListLoan : List of Loan by CustomerID (On Going or Finished)
- GetLoanDetails : Get Loan Details by Loan ID
- GetCustomerData : Get Customer Details

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

## Architecture

### Code Project Structure

Adopting DDD

### Model

#### Loan

```
CREATE TABLE loan (
    ID INTEGER PRIMARY KEY,
    customer_id INTEGER,
    name TEXT,
    amount REAL,
    amount_interest REAL,
    annual_rate_precentage REAL,
    start_date DATE,
    end_date DATE,
    status INTEGER,
    FOREIGN KEY (customer_id) REFERENCES customer(customer_id)
);
```

#### LoanDetail

```
CREATE TABLE loan_details (
	ID INTEGER PRIMARY KEY
    loan_id INTEGER,
    name TEXT,
    amount REAL,
    status INTEGER,
    start_date DATE,
    end_date DATE,
    payment_id INTEGER,
    FOREIGN KEY (loan_id) REFERENCES loan(ID)
);
```

#### Payment
```
CREATE TABLE payment_history (
    payment_id INTEGER PRIMARY KEY,
    name TEXT,
    amount REAL,
    status TEXT
);
```

#### Customer
```
CREATE TABLE customer (
    customer_id INTEGER PRIMARY KEY,
    full_name TEXT NOT NULL,
    status INTEGER,
    credit_status INTEGER
);
```

