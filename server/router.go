package server

import (
	"github.com/gin-gonic/gin"
	"vault-proxy/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	secret := new(controllers.SecretController)
	router.GET("/secrets", secret.Retrieve)

	return router
}
