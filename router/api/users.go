package api

import (
	"net/http"

	"github.com/er230059/golang-practice/service/user"
	"github.com/gin-gonic/gin"
)

type addUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type updateUserDto struct {
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

func AddUser(c *gin.Context) {
	var u addUserDto
	err := c.BindJSON(&u)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, err := user.Create(u.Name, u.Email, u.Password)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func UpdateUser(c *gin.Context) {
	var u updateUserDto
	err := c.BindJSON(&u)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := c.Get("userId")

	err = user.Update(id.(int), u.Name, u.Password)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
}

func GetUser(c *gin.Context) {
	id, _ := c.Get("userId")

	user, err := user.Find(id.(int))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
	})
}
