package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RefreshToken func
func RefreshToken(c *gin.Context) {
	// refreshToken 만료 체크
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
