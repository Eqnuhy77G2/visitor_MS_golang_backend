package middleware

import (
	"VisitorsManagementSystem/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	Username string
	jwt.StandardClaims
}

func ReleaseToken(user models.Admin) (string,error) {
	expirationTime := time.Now().Add(1*time.Hour)
	claims :=&Claims{
		Username: user.Username,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer: "oceanLearn.tech",
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString , err := token.SignedString(jwtKey)

	if err != nil {
		return "",err
	}

	return tokenString,nil

}