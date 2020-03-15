package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/er230059/golang-practice/service/auth"
	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	idFromParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	token := c.Request.Header.Get("Authorization")
	if s := strings.Split(token, " "); len(s) == 2 {
		token = s[1]
	}

	idFromToken, err := auth.Verify(token)

	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	if idFromParams != idFromToken {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	c.Set("userId", idFromParams)
	c.Next()
}
