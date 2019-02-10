package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SecretController struct{}

func (u SecretController) Retrieve(c *gin.Context) {
	dummySecret := make(map[string]interface{})
	dummySecret["key1"] = "val1"
	dummySecret["key2"] = "val2"

	c.JSON(http.StatusOK, gin.H{"message": "Success.", "secrets": dummySecret})
	return
}
