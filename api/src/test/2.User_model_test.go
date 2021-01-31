package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/cheesecat47/webpractice/api/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	json := `
	{
		"UserID": "pororo",
		"UserEmail": "pororo@abc.com",
		"UserPW": "pw4",
		"UserName": "pororo",
		"UserRole": "NULL"
	}`
	log.Printf("TestCreateUser: Create User: json:\n%s", json)

	err := model.CreateUser(json)
	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
	log.Println("TestCreateUser: Create User")
}

// TestGetUserByID func
func TestGetUserByID(t *testing.T) {
	user, err := model.GetUserByID("jy")
	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
	log.Printf("TestGetUser: user:\n%s", string(user))

	user2, err := model.DecodeUserToJSON(user)
	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
	log.Println("TestGetUser: user2:", user2)
}

// TestGetAllUsers func
func TestGetAllUsers(t *testing.T) {
	users, err := model.GetAllUsers()
	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
	log.Printf("TestGetAllUsers: users:\n%s", string(users))
}

// func TestUpdateUser(t *testing.T) {
// 	var user model.User

// 	err := db.Table("user").Where("user_name = ?", "pororo").Take(&user).Update("UserPW", "modifiedPW")
// 	assert.Equal(t, nil, err.Error, fmt.Errorf("TestUpdateUser: Error: %v", err.Error))

// 	log.Println("TestUpdateUser: user:", user)
// }

// func TestDeleteUser(t *testing.T) {
// 	err := db.Table("user").Where("user_name = ?", "pororo").Delete(&model.User{})
// 	assert.Equal(t, nil, err.Error, fmt.Errorf("TestDeleteUser: Error: %v", err.Error))
// }
