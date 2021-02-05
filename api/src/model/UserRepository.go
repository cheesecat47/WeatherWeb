package model

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// UserRepository map
type userRepository struct {
	UserRepo map[string][]User
}

var (
	instance *userRepository
	once     sync.Once
)

// GetUserRepo func
func GetUserRepo() *userRepository {
	once.Do(func() {
		instance = &userRepository{}
		if instance.UserRepo == nil {
			instance.UserRepo = make(map[string][]User)
		}
	})
	return instance
}

func init() {
	log.Println("UserRepository: init")
	instance = GetUserRepo()
}

// CreateUser func
func CreateUser(json string) error {
	u, err := DecodeJSONToUser([]byte(json))
	if err != nil {
		log.Println("UserRepository: CreateUser: err:", err)
		return errorCantCreateUser
	}

	result, err := db.Exec(
		"INSERT INTO user VALUES (?, ?, ?, ?, ?, ?, ?, NULL)",
		u.UserID, u.UserEmail, u.UserPW, u.UserName,
		u.UserRole, time.Now(), time.Now())
	if err != nil {
		log.Println("UserRepository: CreateUser: err:", err)
		return errorCantCreateUser
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}

	log.Println("UserRepository: CreateUser: success")
	return nil
}

// GetUser func
func GetUser(params map[string]string) ([]User, error) {
	query := makeSelectQuery(params)

	val, exists := instance.UserRepo[query]
	if exists {
		log.Println("UserRepository: GetUser: Hit repository")
		return val, nil
	}
	log.Println("UserRepository: GetUser: From querying db")
	result, err := getUserInternal(query)
	instance.UserRepo[query] = result
	return result, err
}

// getUserInternal func
func getUserInternal(query string) ([]User, error) {
	rows, err := db.Query(query)
	if err != nil {
		log.Println("UserRepository: GetUserInternal: err: ", err)
		return nil, err
	}
	defer rows.Close()

	userArray := []User{}
	for rows.Next() {
		temp := RawUser{}
		err = rows.Scan(
			&temp.user.UserID, &temp.user.UserEmail, &temp.user.UserPW, &temp.user.UserName,
			&temp.RawUserRole, &temp.user.CreatedAt, &temp.user.UpdatedAt, &temp.user.DeletedAt)
		if err != nil {
			log.Println("UserRepository: GetUserInternal: err: ", err)
			return nil, err
		}
		userArray = append(userArray, temp.user)
	}

	log.Println("UserRepository: GetUserInternal: success")
	return userArray, nil
}

// GetUserByID func
func GetUserByID(id string) ([]User, error) {
	query := makeSelectQuery(
		map[string]string{
			"from":  "user",
			"where": fmt.Sprintf("user_id=%s", id)})

	val, exists := instance.UserRepo[query]
	if exists {
		log.Println("UserRepository: GetUserByID: Hit repository")
		return val, nil
	}
	log.Println("UserRepository: GetUserByID: From querying db")
	result, err := getUserInternal(query)
	instance.UserRepo[query] = result
	return result, err
}

// GetAllUsers func
func GetAllUsers() ([]byte, error) {

	query := fmt.Sprintf("SELECT * FROM user ORDER BY user_id")
	log.Println("UserRepository: GetAllUsers: query:", query)

	rows, err := db.Query(query)
	if err != nil {
		log.Println("UserRepository: GetAllUsers: err: ", err)
		return nil, err
	}
	defer rows.Close()

	userArray := []User{}
	for rows.Next() {
		ru := RawUser{}
		err := rows.Scan(
			&ru.user.UserID, &ru.user.UserEmail, &ru.user.UserPW, &ru.user.UserName,
			&ru.RawUserRole, &ru.user.CreatedAt, &ru.user.UpdatedAt, &ru.user.DeletedAt)
		if err != nil {
			log.Println("UserRepository: GetAllUsers: err: ", err)
			return nil, err
		}
		ru.checkNullString()
		userArray = append(userArray, ru.user)
	}
	log.Println("UserRepository: GetAllUsers: userArray: ", userArray)

	log.Println("UserRepository: GetAllUsers: success")
	return EncodeUserToJSON(userArray)
}
