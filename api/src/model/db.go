package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/cheesecat47/webpractice/constant"
	_ "github.com/go-sql-driver/mysql"
)

var (
	pool *sql.DB // Database connection pool.
)

var (
	errorCantConnectDB = errors.New("Can't connect DB")
	errorCantOpenDB    = errors.New("Can't open DB")
)

// InitDB func
func InitDB() {
	pool, err := ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}
	if pool != nil {
		pool.SetConnMaxLifetime(0)
		pool.SetMaxIdleConns(3)
		pool.SetMaxOpenConns(3)
	}
	defer pool.Close()
	log.Println("InitDB: pool:", pool)
}

// ConnectDB func
func ConnectDB() (*sql.DB, error) {
	pool, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s",
			constant.MysqlUser,
			constant.MysqlRootPw,
			constant.MysqlHost,
			constant.MysqlDb))
	if err != nil {
		fmt.Println("Database Connection Error")
		return nil, err
	}

	if pool.Ping() != nil {
		fmt.Println("Database Ping Fail")
		return nil, err
	}
	return pool, nil
}

// http://golang.site/go/article/107-MySql-사용---쿼리
// https://golang.org/pkg/database/sql/
// https://bourbonkk.tistory.com/59
// https://pkg.go.dev/github.com/go-sql-driver/mysql?utm_source=gopls#readme-installation
// https://www.popit.kr/golang-databasesql-패키지-삽질기-3편-커넥션-풀/
