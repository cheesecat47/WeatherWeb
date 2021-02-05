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
	db  *sql.DB
	err error
)

var (
	errorCantConnectDB     = errors.New("Can't connect DB")
	errorCantOpenDB        = errors.New("Can't open DB")
	errorCantCloseDB       = errors.New("Can't close DB")
	errorCantEncodeJSON    = errors.New("Can't encode to JSON")
	errorCantDecodeJSON    = errors.New("Can't decode to JSON")
	errorCantCreateUser    = errors.New("Can't create a new user")
	errorQueryParamMissing = errors.New("There's no query parameter")
)

// InitDB func
func InitDB() error {
	log.Println("db.go: InitDB")
	db, err = ConnectDB(
		constant.MysqlUser,
		constant.MysqlRootPw,
		constant.MysqlHost,
		constant.MysqlDb)
	if err != nil {
		log.Println("db.go: InitDB: db: err:", err)
		return err
	}
	if db != nil {
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
		log.Println("db.go: InitDB: Set up db")
	}
	log.Println("db.go: InitDB: db:", db)
	return nil
}

// ConnectDB func
func ConnectDB(user string, pw string, host string, dbname string) (*sql.DB, error) {
	params := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?%s", user, pw, host, dbname, params)
	log.Println("db.go: ConnectDB: dsn:", dsn)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("db.go: ConnectDB: err:", err)
		return nil, errorCantConnectDB
	}
	log.Println("db.go: ConnectDB: db instance is generated")
	return db, nil
}

// CloseDB func
func CloseDB() error {
	err = db.Close()
	if err != nil {
		log.Println("db.go: CloseDB: err:", err)
		return errorCantCloseDB
	}
	log.Println("db.go: CloseDB: db instance is closed")
	return nil
}
