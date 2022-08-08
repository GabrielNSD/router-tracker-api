package router

import (
	"api/controllers/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.POST("/register", users.HandleRegister)
	r.POST("/signin", users.HandleSignin)
	return r
}
