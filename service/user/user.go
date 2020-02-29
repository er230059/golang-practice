package user

import (
	"log"

	"github.com/er230059/golang-practice/database"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Id    int
	Name  string
	Email string
}

func Create(name, email, password string) (int, error) {
	db := database.GetDb()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	hashedPassword := string(hash)

	res, err := db.Exec("INSERT INTO `users`(`name`, `email`, `password`) VALUES (?, ?, ?)", name, email, hashedPassword)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return int(id), nil
}

func Find(id int) (user, error) {
	db := database.GetDb()

	result := user{Id: id}
	err := db.QueryRow("SELECT `name`, `email` from `users` where `id` = ?", id).Scan(&result.Name, &result.Email)
	if err != nil {
		log.Println(err)
		return result, err
	}

	return result, nil
}
