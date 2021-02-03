package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cheesecat47/webpractice/api/controller"
	"github.com/cheesecat47/webpractice/api/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var testUser = model.User{
	UserID:    "1",
	UserEmail: "atg0831@naver.com",
	UserPW:    "1234",
	UserName:  "atg",
}

func TestGenerateToken(t *testing.T) {
	jwtInfo := controller.JwtInfo{
		JwtKey: "SECRETE_KEY",
		Issuer: "github.com/cheesecat47/webpractice",
	}
	generatedToken, err := jwtInfo.GenerateToken(testUser.UserID, testUser.UserEmail)
	assert.NoError(t, err)

	os.Setenv("testToken", generatedToken)
}

func TestValidateToken(t *testing.T) {
	getTestToken := os.Getenv("testToken")
	jwtInfo := controller.JwtInfo{
		JwtKey: "SECRETE_KEY",
		Issuer: "github.com/cheesecat47/webpractice",
	}

	claims, err := jwtInfo.ValidateToken(getTestToken)
	assert.NoError(t, err)
	assert.Equal(t, testUser.UserID, claims.UserID)
	assert.Equal(t, testUser.UserEmail, claims.UserEmail)
}

func TestLogin(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) (c *gin.Context) {
		cTest, _ := gin.CreateTestContext(w)
		cTest.Request = r
		return cTest
	}
	clientUser, err := json.Marshal(testUser)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(clientUser))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c := handler(w, req)

	controller.Login(c)
	assert.Equal(t, 200, w.Code)

}
