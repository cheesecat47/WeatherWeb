package services

import (
	"log"
	"net/http"

	"github.com/cheesecat47/webpractice/api/helper"
	"github.com/cheesecat47/webpractice/api/model"
	"github.com/gin-gonic/gin"
)

//client's temporary user_data for test
var tempUser = model.User{
	UserID:    "1",
	UserEmail: "atg0831",
	UserPW:    "1234",
	UserName:  "atg",
}

//Signup func
func Signup(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()

		return
	}

	//add encrypt password logic

	//insert new userData into user table logic

	c.JSON(http.StatusOK, u)
}

//Login func
func Login(c *gin.Context) {
	var u model.User

	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	//need to get user's data from db
	if "atg0831@naver.com" != u.UserEmail && "1234" != u.UserPW {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "errors": "Wrong Password"})
		return
	}

	//Should set JwtKey secret
	jwtInfo := helper.JwtInfo{AccessKey: "ACCESS_TOKEN", RefreshKey: "REFRESH_KEY", Issuer: "atg0831"}

	//need to change args to user's data from db
	token, err := jwtInfo.GenerateTokenPair("1", "atg0831@naver.com")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnprocessableEntity, "errors": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}
