package api

import (
	"net/http"
	"strconv"

	"github.com/er230059/golang-practice/service/user"
	"github.com/gin-gonic/gin"
)

type addUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

func EditUser(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user, err := user.Find(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
	})
}
