package user

import (
	"log"
	"strings"

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

func Update(id int, name, password *string) error {
	db := database.GetDb()

	query := "UPDATE `users` SET"
	var params []interface{}

	if name != nil {
		query += " `name` = ?,"
		params = append(params, name)
	}

	if password != nil {
		query += " `password` = ?,"

		hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
		}
		hashedPassword := string(hash)

		params = append(params, &hashedPassword)
	}

	if len(params) == 0 {
		return nil
	}

	query = strings.TrimRight(query, ",")
	query += " WHERE `id` = ?"
	params = append(params, &id)

	_, err := db.Exec(query, params...)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
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
