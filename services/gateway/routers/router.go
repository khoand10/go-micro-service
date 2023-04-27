package routers

import (
	"fmt"
	"go-micro-service/services/gateway/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestServer struct {
	Engine *gin.Engine
	Config *config.Config
}

func Init(restServer *RestServer) error {
	apiPrefix := getAPIPrefix(restServer)
	api := restServer.Engine.Group(apiPrefix)

	api.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	InitAuthRouter(api.Group("/auth"))

	port := getPort(restServer)
	err := restServer.Engine.Run(port)
	if err != nil {
		return err
	}
	return nil
}

func getAPIPrefix(server *RestServer) string {
	apiVersion := fmt.Sprintf("/%s", server.Config.RestAPIVersion)
	return fmt.Sprintf("/api%s", apiVersion)
}

func getPort(server *RestServer) string {
	return fmt.Sprintf(":%d", server.Config.RestPort)
}
