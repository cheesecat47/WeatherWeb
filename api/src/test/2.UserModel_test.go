package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/cheesecat47/webpractice/api/model"
	"github.com/stretchr/testify/assert"
)

// func TestCreateUser(t *testing.T) {
// 	json := `
// 	{
// 		"UserID": "pororo",
// 		"UserEmail": "pororo@abc.com",
// 		"UserPW": "pw4",
// 		"UserName": "pororo",
// 		"UserRole": "NULL"
// 	}`
// 	log.Printf("TestCreateUser: Create User: json:\n%s", json)

// 	err := model.CreateUser(json)
// 	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
// 	log.Println("TestCreateUser: Create User")
// }

// TestGetUserByID func
// func TestGetUserByID(t *testing.T) {
// 	user, err := model.GetUserByID("jy")
// 	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
// 	log.Printf("TestGetUser: user:\n%s", string(user))

// 	user2, err := model.DecodeJSONToUser(user)
// 	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
// 	log.Println("TestGetUser: user2:", user2)
// }

// // TestGetAllUsers func
// func TestGetAllUsers(t *testing.T) {
// 	users, err := model.GetAllUsers()
// 	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
// 	log.Printf("TestGetAllUsers: users:\n%s", string(users))
// }

// TestGetUserUsingRepository func
func TestGetUserUsingRepository(t *testing.T) {
	params := map[string]string{
		"from":     "user",
		"order by": "user_name"}
	startTime := time.Now()
	users, err := model.GetUser(params)
	elapsedTime := time.Since(startTime)
	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
	log.Println("TestGetUserUsingRepository: users1:", users)
	log.Println("TestGetUserUsingRepository: elapsedTime:", elapsedTime)
	fmt.Println("")

	startTime = time.Now()
	users, err = model.GetUser(params)
	elapsedTime = time.Since(startTime)
	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
	log.Println("TestGetUserUsingRepository: users2:", users)
	log.Println("TestGetUserUsingRepository: elapsedTime:", elapsedTime)
	fmt.Println("")

	startTime = time.Now()
	users, err = model.GetUser(params)
	elapsedTime = time.Since(startTime)
	assert.Equal(t, nil, err, fmt.Errorf("Error: %v", err))
	log.Println("TestGetUserUsingRepository: users3:", users)
	log.Println("TestGetUserUsingRepository: elapsedTime:", elapsedTime)
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
