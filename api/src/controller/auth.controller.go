package controller

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//JwtInfo struct is secret key for signing token
type JwtInfo struct {
	JwtKey string //"ACCESS_TOKEN"
	Issuer string //"cheesecat47/webpractice"...
}

//Claims struct is for making jwt token
type Claims struct {
	UserID    string
	UserEmail string
	jwt.StandardClaims
}

//GenerateToken func
func (j *JwtInfo) GenerateToken(userID string, userEmail string) (string, error) {
	expTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID:    userID,
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			Issuer:    j.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.JwtKey))
	if err != nil {
		return "", fmt.Errorf("token singed Error")
	}
	return signedToken, nil
}

//GenerateTokenPair func is generating pair token {access_token,refresh_token}
func (j *JwtInfo) GenerateTokenPair(userID string, userEmail string) (map[string]string, error) {
	aTokenExpTime := time.Now().Add(15 * time.Minute)
	aTokenClaims := &Claims{
		UserID:    userID,
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: aTokenExpTime.Unix(),
			Issuer:    j.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aTokenClaims)
	aTokenSigned, err := aToken.SignedString([]byte(j.JwtKey))
	if err != nil {
		return nil, err
	}

	rTokenExpTime := time.Now().Add(time.Hour * 24)
	rTokenClaims := &Claims{
		UserID:    userID,
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: rTokenExpTime.Unix(),
			Issuer:    j.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rTokenClaims)
	rTokenSigned, err := rToken.SignedString([]byte(j.JwtKey))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"accessToken":  aTokenSigned,
		"refreshToken": rTokenSigned,
	}, nil
}

//ValidateToken func
func (j *JwtInfo) ValidateToken(signedToken string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.JwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return nil, err
	}
	return claims, nil

}

//ExtractToken func
func ExtractToken(req string) string {
	clientToken := ""
	bearerToken := req
	// Authorization Bearer xxxx
	strTokenArr := strings.Split(bearerToken, "Bearer ")
	if len(strTokenArr) == 2 {
		clientToken = strings.TrimSpace(strTokenArr[1])
	}

	return clientToken
}

//RefreshToken func
func RefreshToken(c *gin.Context) {

}
