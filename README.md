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

currently no need, credit status customer can be get on the fly for now

## How To Build/Run

### Golang
#### Prerequiested

- Go 1.14 above already installed (tested with 1.23, but should be compatible as long above go 1.11 (go module))
- Git

#### Download the code

ssh `git clone git@github.com:fakihariefnoto/4m4rth4.git`

or

https `git clone https://github.com/fakihariefnoto/4m4rth4.git`

#### Build & Run

Run this command
`make start`

## How To Test

### New Customer Journey

- CreateCustomer

```
curl --location 'http://localhost:9000/api/v1/CreateCustomer' \
--header 'Content-Type: application/json' \
--data '{
    "full_name" : "Fakih Arief Noto"
}'
```

### Existing Customer Journey

- CreateLoan

```
curl --location 'http://localhost:9000/api/v1/CreateLoan' \
--header 'Content-Type: application/json' \
--data '{
    "customer_id" : 1,
    "amount" : 500000
}'
```

- ListLoan
Req by customer id
```
curl --location 'http://localhost:9000/api/v1/ListLoan/1'
```
Resp
```
{
    "data": [
        {
            "id": 1,
            "customer_id": 1,
            "name": "Biaya Pendidikan",
            "amount": 500000,
            "amount_interest": 5000000,
            "rate": 10,
            "start_date": "2024-09-07 14:22:14",
            "end_date": "2025-08-23 14:22:14",
            "status": "OnGoing",
            "details": null,
            "total_paid": 0,
            "total_borrowed": 0,
            "outstanding_balance": 0
        }
    ],
    "error": false,
    "msg": null
}
```

- GetOutstanding
Request
```
curl --location 'http://localhost:9000/api/v1/GetOutstanding/1'
```
Resp
```
{
    "data": {
        "id": 1,
        "customer_id": 1,
        "name": "Biaya Pendidikan",
        "amount": 500000,
        "amount_interest": 5000000,
        "rate": 10,
        "start_date": "2024-09-07 14:22:14",
        "end_date": "2025-08-23 14:22:14",
        "status": "OnGoing",
        "details": [
            {
                "id": 1,
                "loan_id": 1,
                "name": "Week-1",
                "amount": 110000,
                "start_date": "2024-09-07 14:22:14",
                "end_date": "2024-09-14 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 2,
                "loan_id": 1,
                "name": "Week-2",
                "amount": 110000,
                "start_date": "2024-09-14 14:22:14",
                "end_date": "2024-09-21 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 3,
                "loan_id": 1,
                "name": "Week-3",
                "amount": 110000,
                "start_date": "2024-09-21 14:22:14",
                "end_date": "2024-09-28 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 4,
                "loan_id": 1,
                "name": "Week-4",
                "amount": 110000,
                "start_date": "2024-09-28 14:22:14",
                "end_date": "2024-10-05 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
           ....
        ],
        "total_paid": 0,
        "total_borrowed": 5500000,
        "outstanding_balance": 5500000
    },
    "error": false,
    "msg": null
}
```

- IsDelinquent / GetCustomer
```
curl --location 'http://localhost:9000/api/v1/IsDelinquent/1'
```

```
curl --location 'http://localhost:9000/api/v1/GetCustomer/1'
```

- GetLoanDetails
Req. to get detail loan id
```
curl --location 'http://localhost:9000/api/v1/GetLoanDetails/1'
```

- MakePayment
```
curl --location 'http://localhost:9000/api/v1/MakePayment' \
--header 'Content-Type: application/json' \
--data '{
    "customer_id" : 1,
    "loan_detail_id" : 1,
    "amount" : 110000
}'
```
the response after payment form outstanding
```
{
    "data": {
        "id": 1,
        "customer_id": 1,
        "name": "Biaya Pendidikan",
        "amount": 500000,
        "amount_interest": 5000000,
        "rate": 10,
        "start_date": "2024-09-07 14:22:14",
        "end_date": "2025-08-23 14:22:14",
        "status": "OnGoing",
        "details": [
            {
                "id": 1,
                "loan_id": 1,
                "name": "Week-1",
                "amount": 110000,
                "start_date": "2024-09-07 14:22:14",
                "end_date": "2024-09-14 14:22:14",
                "status": "Paid",
                "payment_id": 2
            },
          ...
        ],
        "total_paid": 110000,
        "total_borrowed": 5500000,
        "outstanding_balance": 5390000,
        "OverDueCounter": 0
    },
    "error": false,
    "msg": null
}
```

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

