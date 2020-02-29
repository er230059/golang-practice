package middleware

import (
	"net/http"
	"strings"

	"github.com/er230059/golang-practice/service/auth"
	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if s := strings.Split(token, " "); len(s) == 2 {
		token = s[1]
	}

	id, err := auth.Verify(token)

	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	c.Set("userId", id)
	c.Next()
}
