package api

import (
	"net/http"

	"github.com/er230059/golang-practice/service/auth"
	"github.com/gin-gonic/gin"
)

type loginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var u loginDto
	err := c.BindJSON(&u)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	token, err := auth.Login(u.Email, u.Password)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.Status(http.StatusUnauthorized)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
