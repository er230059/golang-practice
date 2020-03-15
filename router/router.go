package router

import (
	"github.com/er230059/golang-practice/router/api"
	"github.com/er230059/golang-practice/router/middleware"
	"github.com/gin-gonic/gin"
)

//InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.POST("/login", api.Login)
	r.POST("/users", api.AddUser)

	authorized := r.Group("/")
	authorized.Use(middleware.VerifyToken)
	{
		authorized.GET("/users/:id", api.GetUser)
		authorized.PATCH("/users/:id", api.UpdateUser)
	}

	return r
}
