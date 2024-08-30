package database

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/fakihariefnoto/4m4rth4/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

var (
	mapDB         map[ConnectionName]*sql.DB
	mapConnString map[ConnectionName]string
)

type (
	ConnectionName string

	dbConn struct {
		db *sql.DB
	}
)

const (
	CustomerDB ConnectionName = "customer"
	PaymentDB  ConnectionName = "payment"
	LoanDB     ConnectionName = "loan"
)

func init() {
	mapDB = make(map[ConnectionName]*sql.DB)
	mapConnString = make(map[ConnectionName]string)
}

func CreateDB(dbName string) error {
	if _, err := os.Stat(dbName); os.IsExist(err) {
		// err = os.Remove(dbName)
		// if err != nil {
		// 	return err
		// }
		log.Println("CreateDB : DB ", dbName, " Already exist")
		return nil
	}

	log.Println("CreateDB : DB Not exsit, Creating.. ", dbName)

	_, err := os.Create(dbName)
	if err != nil {
		return err
	}

	log.Println("CreateDB :  DB ", dbName, " Created Successfully!")
	return nil
}

func Init(conf config.Config) {
	for _, db := range conf.DB {
		mapConnString[ConnectionName(db.Name)] = db.ConnectionString
	}
}

func AddConnection(dbName ConnectionName) (err error) {

	connString := mapConnString[dbName]

	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return err
	}
	mapDB[dbName] = db
	return
}

func GetDBConn(dbName ConnectionName) (*dbConn, error) {
	if mapDB[dbName] == nil {
		return nil, errors.New("DB not exist")
	}
	return &dbConn{db: mapDB[dbName]}, nil
}

func Disconnect(dbName ConnectionName) error {
	if _, exist := mapDB[dbName]; !exist {
		return errors.New("DB not exist or not connected")
	}
	if mapDB[dbName] == nil {
		return errors.New("DB have no connection")
	}
	return mapDB[dbName].Close()
}

func (d *dbConn) Select(query string, args ...interface{}) (*sql.Rows, error) {
	res, err := d.db.Query(query, args...)
	return res, err
}

func (d *dbConn) Exec(query string, args ...interface{}) (sql.Result, error) {
	res, err := d.db.Exec(query, args...)
	return res, err
}

// func createTable() {
// 	log.Println("createTable user")
// 	syntax := `
// 	CREATE TABLE IF NOT EXISTS user (
// 		user_id INTEGER PRIMARY KEY,
// 		name TEXT NOT NULL,
// 		sex TEXT NOT NULL,
// 		age TEXT NOT NULL
// 	);
// 	`
// 	// phone TEXT NOT NULL UNIQUE
// 	_, err := db.Exec(syntax)
// 	if err != nil {
// 		log.Fatalln("createTable, ", err)
// 	}
// }

/*

CREATE TABLE customer (
    customer_id INTEGER PRIMARY KEY,
    full_name TEXT NOT NULL,
    status INTEGER,
    credit_status INTEGER
);

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

CREATE TABLE payment_history (
    payment_id INTEGER PRIMARY KEY,
    name TEXT,
    amount REAL,
    status TEXT
);

*/
