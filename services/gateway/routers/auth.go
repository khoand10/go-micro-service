package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(g *gin.RouterGroup) {
	g.POST("/login", login())
	g.GET("/login", login())
}

func login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "login test",
		})
	}
}
