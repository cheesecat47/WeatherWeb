package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// User struct
type User struct {
	UserID    string `db:"user_id"`
	UserEmail string `db:"user_email"`
	UserPW    string `db:"user_pw"`
	UserName  string `db:"user_name"`
	UserRole  string
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

// RawUser struct
type RawUser struct {
	user        User
	RawUserRole sql.NullString `db:"user_role"`
}

// checkNullString func
func (ru *RawUser) checkNullString() {
	if ru.RawUserRole.Valid {
		// log.Println("User_model: checkNullString: valid")
		ru.user.UserRole = ru.RawUserRole.String
	} else {
		// log.Println("User_model: checkNullString: NULL")
		ru.user.UserRole = "NULL"
	}
	log.Println("User_model: checkNullString: success")
}

// EncodeUserToJSON func
func EncodeUserToJSON(v interface{}) ([]byte, error) {
	jsonBytes, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		log.Println("User_model: EncodeUserToJSON: err:", err)
		return nil, errorCantEncodeJSON
	}
	log.Println("User_model: EncodeUserToJSON: success")
	return jsonBytes, nil
}

// DecodeUserToJSON func
func DecodeUserToJSON(jsonBytes []byte) (User, error) {
	var temp User
	err := json.Unmarshal(jsonBytes, &temp)
	if err != nil {
		log.Println("User_model: DecodeUserToJSON: err:", err)
		return temp, errorCantDecodeJSON
	}
	log.Println("User_model: DecodeUserToJSON: success")
	return temp, nil
}

// CreateUser func
func CreateUser(json string) error {
	u, err := DecodeUserToJSON([]byte(json))
	if err != nil {
		log.Println("User_model: CreateUser: err:", err)
		return errorCantCreateUser
	}

	result, err := db.Exec(
		"INSERT INTO user VALUES (?, ?, ?, ?, ?, ?, ?, NULL)",
		u.UserID, u.UserEmail, u.UserPW, u.UserName,
		u.UserRole, time.Now(), time.Now())
	if err != nil {
		log.Println("User_model: CreateUser: err:", err)
		return errorCantCreateUser
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}

	log.Println("User_model: CreateUser: success")
	return nil
}

// GetUserByID func
func GetUserByID(id string) ([]byte, error) {
	ru := RawUser{}
	query := fmt.Sprintf("SELECT * FROM user WHERE user_id = '%s'", id)
	// log.Println("User_model: GetUserByID: query:", query)

	err := db.QueryRow(query).Scan(
		&ru.user.UserID, &ru.user.UserEmail, &ru.user.UserPW, &ru.user.UserName,
		&ru.RawUserRole, &ru.user.CreatedAt, &ru.user.UpdatedAt, &ru.user.DeletedAt)
	if err != nil {
		log.Println("User_model: GetUserByID: err: ", err)
		return nil, err
	}

	ru.checkNullString()
	log.Println("User_model: GetUserByID: success")
	return EncodeUserToJSON(ru.user)
}

// GetAllUsers func
func GetAllUsers() ([]byte, error) {

	query := fmt.Sprintf("SELECT * FROM user ORDER BY user_id")
	log.Println("User_model: GetAllUsers: query:", query)

	rows, err := db.Query(query)
	// .Scan(
	// 	&u.UserID, &u.UserEmail, &u.UserPW, &u.UserName,
	// 	&u.RawUserRole, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		log.Println("User_model: GetAllUsers: err: ", err)
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
			log.Println("User_model: GetAllUsers: err: ", err)
			return nil, err
		}
		ru.checkNullString()
		userArray = append(userArray, ru.user)
	}
	log.Println("User_model: GetAllUsers: userArray: ", userArray)

	log.Println("User_model: GetAllUsers: success")
	return EncodeUserToJSON(userArray)
}
