package test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cheesecat47/webpractice/api/constant"
	"github.com/cheesecat47/webpractice/api/model"
	"github.com/stretchr/testify/assert"
)

var (
	db  *sql.DB
	err error
)

// TestDBConnection func
func TestDBConnection(t *testing.T) {
	constant.MysqlUser = "root"
	constant.MysqlRootPw = "1234"
	constant.MysqlHost = "localhost:33010"
	constant.MysqlDb = "webpracticedb"
	err = model.InitDB()
	assert.Equal(t, nil, err, fmt.Errorf("TestDBConnection: Error: %v", err))
}

// TestClosedDB func
// func TestClosedDB(t *testing.T) {
// 	err = model.CloseDB()
// 	assert.Equal(t, nil, err,
// 		fmt.Errorf("TestClosedDB: Error: %v", err))
// }
