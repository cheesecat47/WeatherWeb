package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cheesecat47/webpractice/api/constant"
	// mariadb driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB Database connection pool.
	DB  *sql.DB
	err error
)

var (
	errorCantConnectDB = errors.New("Can't connect DB")
	errorCantOpenDB    = errors.New("Can't open DB")
	errorCantCloseDB   = errors.New("Can't close DB")
)

// InitDB func
func InitDB() error {
	log.Println("db.go: InitDB")
	DB, err = ConnectDB(
		constant.MysqlUser,
		constant.MysqlRootPw,
		constant.MysqlHost,
		constant.MysqlDb)
	if err != nil {
		log.Println("db.go: InitDB: db: err != nil")
		return err
	}
	if DB != nil {
		DB.SetConnMaxLifetime(time.Minute * 3)
		DB.SetMaxIdleConns(10)
		DB.SetMaxOpenConns(10)
		log.Println("db.go: InitDB: Set up db")
	}
	log.Println("db.go: InitDB: db:", DB)
	return nil
}

// ConnectDB func
func ConnectDB(user string, pw string, host string, dbname string) (*sql.DB, error) {
	params := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?%s", user, pw, host, dbname, params)
	log.Println("db.go: ConnectDB: dsn:", dsn)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("db.go: ConnectDB: err != nil")
		return nil, errorCantConnectDB
	}
	log.Println("db.go: ConnectDB: db instance is generated")
	return DB, nil
}

// CloseDB func
func CloseDB() error {
	log.Println("db.go: CloseDB")
	err = DB.Close()
	if err != nil {
		return errorCantCloseDB
	}
	log.Println("db.go: CloseDB: db instance is closed")
	return nil
}
