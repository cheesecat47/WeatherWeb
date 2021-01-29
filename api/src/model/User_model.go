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
	User
	RawUserRole sql.NullString `db:"user_role"`
}

// checkNullString func
func (u *RawUser) checkNullString() {
	if u.RawUserRole.Valid {
		log.Println("User_model: checkNullString: valid")
		u.UserRole = u.RawUserRole.String
	} else {
		log.Println("User_model: checkNullString: NULL")
		u.UserRole = "NULL"
	}
}

// EncodeUserToJSON func
func (u *User) EncodeUserToJSON() ([]byte, error) {
	jsonBytes, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		log.Println("User_model: EncodeUserToJSON: err:", err)
		return nil, errorCantEncodeJSON
	}
	log.Println("User_model: EncodeUserToJSON: success")
	return jsonBytes, nil
}

// DecodeUserToJSON func
func DecodeUserToJSON(jsonBytes []byte) (User, error) {
	var user User
	err := json.Unmarshal(jsonBytes, &user)
	if err != nil {
		log.Println("User_model: DecodeUserToJSON: err:", err)
		return user, errorCantDecodeJSON
	}
	log.Println("User_model: DecodeUserToJSON: success")
	return user, nil
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
	u := RawUser{}
	query := fmt.Sprintf("SELECT * FROM user WHERE user_id = '%s'", id)
	// log.Println("User_model: GetUserByID: query:", query)

	err := db.QueryRow(query).Scan(
		&u.UserID, &u.UserEmail, &u.UserPW, &u.UserName,
		&u.RawUserRole, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		log.Println("User_model: GetUserByID: err: ", err)
		return nil, err
	}

	u.checkNullString()
	log.Println("User_model: GetUserByID: success")
	return u.EncodeUserToJSON()
}
