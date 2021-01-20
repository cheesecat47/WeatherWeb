package constant

import (
	"os"
)

var (
	ListeningPort string = os.Getenv("PORT")
	MysqlDb       string = os.Getenv("MYSQL_DATABASE")
	MysqlHost     string = os.Getenv("MYSQL_HOST")
	MysqlRootPw   string = os.Getenv("MYSQL_ROOT_PASSWORD")
	MysqlUser     string = os.Getenv("MYSQL_USER")
)
