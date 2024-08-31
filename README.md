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
            {
                "id": 5,
                "loan_id": 1,
                "name": "Week-5",
                "amount": 110000,
                "start_date": "2024-10-05 14:22:14",
                "end_date": "2024-10-12 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 6,
                "loan_id": 1,
                "name": "Week-6",
                "amount": 110000,
                "start_date": "2024-10-12 14:22:14",
                "end_date": "2024-10-19 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 7,
                "loan_id": 1,
                "name": "Week-7",
                "amount": 110000,
                "start_date": "2024-10-19 14:22:14",
                "end_date": "2024-10-26 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 8,
                "loan_id": 1,
                "name": "Week-8",
                "amount": 110000,
                "start_date": "2024-10-26 14:22:14",
                "end_date": "2024-11-02 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 9,
                "loan_id": 1,
                "name": "Week-9",
                "amount": 110000,
                "start_date": "2024-11-02 14:22:14",
                "end_date": "2024-11-09 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 10,
                "loan_id": 1,
                "name": "Week-10",
                "amount": 110000,
                "start_date": "2024-11-09 14:22:14",
                "end_date": "2024-11-16 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 11,
                "loan_id": 1,
                "name": "Week-11",
                "amount": 110000,
                "start_date": "2024-11-16 14:22:14",
                "end_date": "2024-11-23 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 12,
                "loan_id": 1,
                "name": "Week-12",
                "amount": 110000,
                "start_date": "2024-11-23 14:22:14",
                "end_date": "2024-11-30 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 13,
                "loan_id": 1,
                "name": "Week-13",
                "amount": 110000,
                "start_date": "2024-11-30 14:22:14",
                "end_date": "2024-12-07 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 14,
                "loan_id": 1,
                "name": "Week-14",
                "amount": 110000,
                "start_date": "2024-12-07 14:22:14",
                "end_date": "2024-12-14 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 15,
                "loan_id": 1,
                "name": "Week-15",
                "amount": 110000,
                "start_date": "2024-12-14 14:22:14",
                "end_date": "2024-12-21 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 16,
                "loan_id": 1,
                "name": "Week-16",
                "amount": 110000,
                "start_date": "2024-12-21 14:22:14",
                "end_date": "2024-12-28 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 17,
                "loan_id": 1,
                "name": "Week-17",
                "amount": 110000,
                "start_date": "2024-12-28 14:22:14",
                "end_date": "2025-01-04 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 18,
                "loan_id": 1,
                "name": "Week-18",
                "amount": 110000,
                "start_date": "2025-01-04 14:22:14",
                "end_date": "2025-01-11 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 19,
                "loan_id": 1,
                "name": "Week-19",
                "amount": 110000,
                "start_date": "2025-01-11 14:22:14",
                "end_date": "2025-01-18 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 20,
                "loan_id": 1,
                "name": "Week-20",
                "amount": 110000,
                "start_date": "2025-01-18 14:22:14",
                "end_date": "2025-01-25 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 21,
                "loan_id": 1,
                "name": "Week-21",
                "amount": 110000,
                "start_date": "2025-01-25 14:22:14",
                "end_date": "2025-02-01 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 22,
                "loan_id": 1,
                "name": "Week-22",
                "amount": 110000,
                "start_date": "2025-02-01 14:22:14",
                "end_date": "2025-02-08 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 23,
                "loan_id": 1,
                "name": "Week-23",
                "amount": 110000,
                "start_date": "2025-02-08 14:22:14",
                "end_date": "2025-02-15 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 24,
                "loan_id": 1,
                "name": "Week-24",
                "amount": 110000,
                "start_date": "2025-02-15 14:22:14",
                "end_date": "2025-02-22 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 25,
                "loan_id": 1,
                "name": "Week-25",
                "amount": 110000,
                "start_date": "2025-02-22 14:22:14",
                "end_date": "2025-03-01 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 26,
                "loan_id": 1,
                "name": "Week-26",
                "amount": 110000,
                "start_date": "2025-03-01 14:22:14",
                "end_date": "2025-03-08 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 27,
                "loan_id": 1,
                "name": "Week-27",
                "amount": 110000,
                "start_date": "2025-03-08 14:22:14",
                "end_date": "2025-03-15 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 28,
                "loan_id": 1,
                "name": "Week-28",
                "amount": 110000,
                "start_date": "2025-03-15 14:22:14",
                "end_date": "2025-03-22 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 29,
                "loan_id": 1,
                "name": "Week-29",
                "amount": 110000,
                "start_date": "2025-03-22 14:22:14",
                "end_date": "2025-03-29 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 30,
                "loan_id": 1,
                "name": "Week-30",
                "amount": 110000,
                "start_date": "2025-03-29 14:22:14",
                "end_date": "2025-04-05 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 31,
                "loan_id": 1,
                "name": "Week-31",
                "amount": 110000,
                "start_date": "2025-04-05 14:22:14",
                "end_date": "2025-04-12 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 32,
                "loan_id": 1,
                "name": "Week-32",
                "amount": 110000,
                "start_date": "2025-04-12 14:22:14",
                "end_date": "2025-04-19 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 33,
                "loan_id": 1,
                "name": "Week-33",
                "amount": 110000,
                "start_date": "2025-04-19 14:22:14",
                "end_date": "2025-04-26 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 34,
                "loan_id": 1,
                "name": "Week-34",
                "amount": 110000,
                "start_date": "2025-04-26 14:22:14",
                "end_date": "2025-05-03 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 35,
                "loan_id": 1,
                "name": "Week-35",
                "amount": 110000,
                "start_date": "2025-05-03 14:22:14",
                "end_date": "2025-05-10 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 36,
                "loan_id": 1,
                "name": "Week-36",
                "amount": 110000,
                "start_date": "2025-05-10 14:22:14",
                "end_date": "2025-05-17 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 37,
                "loan_id": 1,
                "name": "Week-37",
                "amount": 110000,
                "start_date": "2025-05-17 14:22:14",
                "end_date": "2025-05-24 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 38,
                "loan_id": 1,
                "name": "Week-38",
                "amount": 110000,
                "start_date": "2025-05-24 14:22:14",
                "end_date": "2025-05-31 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 39,
                "loan_id": 1,
                "name": "Week-39",
                "amount": 110000,
                "start_date": "2025-05-31 14:22:14",
                "end_date": "2025-06-07 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 40,
                "loan_id": 1,
                "name": "Week-40",
                "amount": 110000,
                "start_date": "2025-06-07 14:22:14",
                "end_date": "2025-06-14 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 41,
                "loan_id": 1,
                "name": "Week-41",
                "amount": 110000,
                "start_date": "2025-06-14 14:22:14",
                "end_date": "2025-06-21 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 42,
                "loan_id": 1,
                "name": "Week-42",
                "amount": 110000,
                "start_date": "2025-06-21 14:22:14",
                "end_date": "2025-06-28 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 43,
                "loan_id": 1,
                "name": "Week-43",
                "amount": 110000,
                "start_date": "2025-06-28 14:22:14",
                "end_date": "2025-07-05 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 44,
                "loan_id": 1,
                "name": "Week-44",
                "amount": 110000,
                "start_date": "2025-07-05 14:22:14",
                "end_date": "2025-07-12 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 45,
                "loan_id": 1,
                "name": "Week-45",
                "amount": 110000,
                "start_date": "2025-07-12 14:22:14",
                "end_date": "2025-07-19 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 46,
                "loan_id": 1,
                "name": "Week-46",
                "amount": 110000,
                "start_date": "2025-07-19 14:22:14",
                "end_date": "2025-07-26 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 47,
                "loan_id": 1,
                "name": "Week-47",
                "amount": 110000,
                "start_date": "2025-07-26 14:22:14",
                "end_date": "2025-08-02 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 48,
                "loan_id": 1,
                "name": "Week-48",
                "amount": 110000,
                "start_date": "2025-08-02 14:22:14",
                "end_date": "2025-08-09 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 49,
                "loan_id": 1,
                "name": "Week-49",
                "amount": 110000,
                "start_date": "2025-08-09 14:22:14",
                "end_date": "2025-08-16 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            },
            {
                "id": 50,
                "loan_id": 1,
                "name": "Week-50",
                "amount": 110000,
                "start_date": "2025-08-16 14:22:14",
                "end_date": "2025-08-23 14:22:14",
                "status": "NotDue",
                "payment_id": 0
            }
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

