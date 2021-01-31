package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var UserRepository = make(map[string]User)

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
// func GetUser() (*User, error) {
// 	if user in UserRepository {
// 		log.Println("UserRepository: GetUser: Hit cache")

// 	} else {
// 		log.Println("UserRepository: GetUser: From querying db")
// 		getUserInternal()
// 	}
// }

// getUserInternal func
func GetUserInternal(params map[string]string) ([]User, error) {
	query := fmt.Sprintf("SELECT")
	paramOrder := [...]string{"attr", "from", "where", "order by", "limit"}

	for _, p := range paramOrder {
		switch val, exists := params[p]; p {
		case "attr":
			if !exists {
				val = "*"
			}
			query += fmt.Sprintf(" %s", val)
		case "from":
			if !exists {
				return nil, errorQueryParamMissing
			}
			query += fmt.Sprintf(" %s %s", p, val)
		case "where", "order by", "limit":
			if !exists {
				continue
			}
			query += fmt.Sprintf(" %s %s", p, val)
		}
	}
	log.Println("UserRepository: GetUserInternal: query:", query)

	rows, err := db.Query(query)
	if err != nil {
		log.Println("UserRepository: GetUserInternal: err: ", err)
		return nil, err
	}
	defer rows.Close()

	userArray := []User{}
	for rows.Next() {
		ru := RawUser{}
		_, exists := params["attr"]
		if !exists {
			err = rows.Scan(
				&ru.user.UserID, &ru.user.UserEmail, &ru.user.UserPW, &ru.user.UserName,
				&ru.RawUserRole, &ru.user.CreatedAt, &ru.user.UpdatedAt, &ru.user.DeletedAt)
		} else {
			err = MyScan(rows, &ru)
		}
		if err != nil {
			log.Println("UserRepository: GetUserInternal: err: ", err)
			return nil, err
		}
		ru.checkNullString()
		userArray = append(userArray, ru.user)
	}
	log.Println("UserRepository: GetUserInternal: userArray: ", userArray)

	log.Println("UserRepository: GetUserInternal: success")
	return userArray, nil
}

// GetUserByID func
func GetUserByID(id string) ([]byte, error) {
	ru := RawUser{}
	query := fmt.Sprintf("SELECT * FROM user WHERE user_id = '%s'", id)
	// log.Println("User_model: GetUserByID: query:", query)

	row := db.QueryRow(query)
	err := row.Scan(
		&ru.user.UserID, &ru.user.UserEmail, &ru.user.UserPW, &ru.user.UserName,
		&ru.RawUserRole, &ru.user.CreatedAt, &ru.user.UpdatedAt, &ru.user.DeletedAt)
	if err != nil {
		log.Println("UserRepository: GetUserByID: err: ", err)
		return nil, err
	}

	ru.checkNullString()
	log.Println("UserRepository: GetUserByID: success")
	return EncodeUserToJSON(ru.user)
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

// MyScan func
func MyScan(rs *sql.Rows, dest interface{}) error {
	log.Println("UserRepository: MyScan: rs:", rs)

	log.Println("UserRepository: MyScan: dest:", dest)
	return nil
}
