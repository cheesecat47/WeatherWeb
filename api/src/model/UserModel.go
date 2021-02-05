package model

import (
	"database/sql"
	"encoding/json"
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

// DecodeJSONToUser func
func DecodeJSONToUser(jsonBytes []byte) (User, error) {
	var temp User
	err := json.Unmarshal(jsonBytes, &temp)
	if err != nil {
		log.Println("User_model: DecodeJSONToUser: err:", err)
		return temp, errorCantDecodeJSON
	}
	log.Println("User_model: DecodeJSONToUser: success")
	return temp, nil
}

