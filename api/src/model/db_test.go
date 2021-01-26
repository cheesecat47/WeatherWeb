package model

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	db, err := ConnectDB("root", "1234", "localhost:33010", "webpracticedb")
	assert.EqualValues(t, "*gorm.DB", fmt.Sprintf("%T", db), "TestDBConnection: DB type err")
	assert.Equal(t, nil, err, fmt.Errorf("TestDBConnection: Error: %v", err))
}

func TestCreateUser(t *testing.T) {
	db, _ := ConnectDB("root", "1234", "localhost:33010", "webpracticedb")

	user := User{
		UserID:    "pororo",
		UserEmail: "pororo@abc.com",
		UserPW:    "pw4",
		UserName:  "pororo"}
	err := db.Table("user").Create(&user)
	assert.Equal(t, nil, err.Error, fmt.Errorf("Error: %v", err.Error))
	log.Println("TestCreateUser: Create User")
}

func TestGetUser(t *testing.T) {
	db, _ := ConnectDB("root", "1234", "localhost:33010", "webpracticedb")
	var user User

	err := db.Table("user").Where("user_name = ?", "pororo").Take(&user)
	assert.Equal(t, nil, err.Error, fmt.Errorf("TestGetUser: Error: %v", err.Error))

	log.Println("TestGetUser: user:", user)
}

func TestUpdateUser(t *testing.T) {
	db, _ := ConnectDB("root", "1234", "localhost:33010", "webpracticedb")
	var user User

	err := db.Table("user").Where("user_name = ?", "pororo").Take(&user).Update("UserPW", "modifiedPW")
	assert.Equal(t, nil, err.Error, fmt.Errorf("TestUpdateUser: Error: %v", err.Error))

	log.Println("TestUpdateUser: user:", user)
}

func TestDeleteUser(t *testing.T) {
	db, _ := ConnectDB("root", "1234", "localhost:33010", "webpracticedb")

	err := db.Table("user").Where("user_name = ?", "pororo").Delete(&User{})
	assert.Equal(t, nil, err.Error, fmt.Errorf("TestDeleteUser: Error: %v", err.Error))
}
