package controller

import (
	"log"
	"net/http"

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

	//Get user data from model?...
	if tempUser.UserName != u.UserName && tempUser.UserPW != u.UserPW {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "errors": "give me vaild login info"})
		return
	}
	//Should set JwtKey secret
	jwtInfo := JwtInfo{JwtKey: "ACCESS_TOKEN", Issuer: "atg0831"}
	token, err := jwtInfo.GenerateToken(tempUser.UserID, tempUser.UserEmail)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnprocessableEntity, "errors": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}

//Logout func
func Logout(c *gin.Context) {

}

//DeleteUser func delete user data
func DeleteUser(c *gin.Context) {
	claim, _ := c.Get("userInfo")

	log.Println(claim)
	c.JSON(http.StatusOK, claim)

	//Delete user data from DB logic

}

//modify user's password data func
