package database

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	mapDB map[ConnectionName]*sql.DB
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
}

func CreateDB(dbName string) error {
	if _, err := os.Stat(dbName); os.IsExist(err) {
		err = os.Remove(dbName)
		if err != nil {
			return err
		}
	}
	_, err := os.Create(dbName)
	if err != nil {
		return err
	}
	return nil
}

func Init()

func AddConnection(dbName ConnectionName) (err error) {
	db, err := sql.Open("sqlite3", string(dbName))
	if err != nil {
		log.Fatal("connect, ", err)
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
