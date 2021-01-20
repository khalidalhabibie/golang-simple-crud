package handler

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userid uint, username string) string {
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
		"iat":      time.Now().Unix(),
		"userID":   userid,
		"userName": username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t

}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func ParsingTokentoUsername(token string) string {
	claims := jwt.MapClaims{}
	result, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil

	})

	if err != nil {
		fmt.Println(err)
	}
	log.Print(result)
	return claims["userName"].(string)

}

func ParsingTokentoID(token string) int {
	claims := jwt.MapClaims{}
	result, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil

	})

	if err != nil {
		fmt.Println(err)
	}
	log.Print(result)

	return int(claims["userID"].(float64))

}
