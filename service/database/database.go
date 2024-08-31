package database

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"billingapp/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

var (
	mapDB         map[ConnectionName]*sql.DB
	mapConnString map[ConnectionName]string
)

type (
	ConnectionName string

	dbConn struct {
		DB *sql.DB
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
	if _, err := os.Stat(dbName); os.IsExist(err) || err == nil {
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
	return &dbConn{DB: mapDB[dbName]}, nil
}

func Connect(dbName string) (db *sql.DB, err error) {
	log.Println("connect ", dbName)
	db, err = sql.Open("sqlite3", dbName)

	if err != nil {
		log.Fatal("connect 2, ", err)
	}
	return
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
	res, err := d.DB.Query(query, args...)
	return res, err
}

func (d *dbConn) Exec(query string, args ...interface{}) (sql.Result, error) {
	res, err := d.DB.Exec(query, args...)
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




 */
