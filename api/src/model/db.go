package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/cheesecat47/webpractice/api/constant"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB // Database connection pool.
)

var (
	errorCantConnectDB = errors.New("Can't connect DB")
	errorCantOpenDB    = errors.New("Can't open DB")
)

// InitDB func
func InitDB() {
	db, err := ConnectDB(
		constant.MysqlUser,
		constant.MysqlRootPw,
		constant.MysqlHost,
		constant.MysqlDb)
	if err != nil {
		log.Println("db.go: InitDB: db: err != nil")
	}
	// if db != nil {
	// 	db.SetConnMaxLifetime(time.Minute * 3)
	// 	db.SetMaxIdleConns(3)
	// 	db.SetMaxOpenConns(3)
	// 	log.Println("db.go: InitDB: Set up db")
	// }
	log.Println("db.go: InitDB: db:", db)
}

// ConnectDB func
func ConnectDB(user string, pw string, host string, dbname string) (*gorm.DB, error) {
	// time format: https://qiita.com/shaching/items/6ab2e3bbc9c2a974eecf
	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?%s", user, pw, host, dbname, param)
	log.Println("db.go: ConnectDB: dsn:", dsn)

	// db, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("db.go: ConnectDB: err != nil")
		return nil, errorCantConnectDB
	}
	log.Println("db.go: ConnectDB: db instance is generated")
	return db, nil
}

// http://golang.site/go/article/107-MySql-사용---쿼리
// https://golang.org/pkg/database/sql/
// https://bourbonkk.tistory.com/59
// https://pkg.go.dev/github.com/go-sql-driver/mysql?utm_source=gopls#readme-installation
// https://www.popit.kr/golang-databasesql-패키지-삽질기-3편-커넥션-풀/
// https://github.com/go-gorp/gorp
