package auth

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/er230059/golang-practice/database"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (string, error) {
	db := database.GetDb()

	var id int
	var hashedPassword string

	err := db.QueryRow("SELECT `id`, `password` from `users` where `email` = ? ", email).Scan(&id, &hashedPassword)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

func Verify(tokenString string) (id int, err error) {
	type CustomClaims struct {
		Id int `json:"id"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return 0, errors.New("Invalid token")
	}

	return claims.Id, nil
}
