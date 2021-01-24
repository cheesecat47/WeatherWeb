package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cheesecat47/webpractice/api/constant"
	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	pool *sql.DB // Database connection pool.
// )

var (
	errorCantConnectDB = errors.New("Can't connect DB")
	errorCantOpenDB    = errors.New("Can't open DB")
)

// InitDB func
func InitDB() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	if db != nil {
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxIdleConns(3)
		db.SetMaxOpenConns(3)
		log.Println("InitDB: Set up db")
	}
	defer db.Close()
	log.Println("InitDB: db:", db)

	// for testing connection
	var (
		id   string
		name string
	)
	rows, err := db.Query("select user_id, user_name from user where user_id = ?", "jy")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

// ConnectDB func
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@(%s)/%s",
			constant.MysqlUser,
			constant.MysqlRootPw,
			constant.MysqlHost,
			constant.MysqlDb))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// http://golang.site/go/article/107-MySql-사용---쿼리
// https://golang.org/pkg/database/sql/
// https://bourbonkk.tistory.com/59
// https://pkg.go.dev/github.com/go-sql-driver/mysql?utm_source=gopls#readme-installation
// https://www.popit.kr/golang-databasesql-패키지-삽질기-3편-커넥션-풀/
