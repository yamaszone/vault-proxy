package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"vault-proxy/server"
	"vault-proxy/util"
)

func TestRootEndpoint(t *testing.T) {
	router := server.NewRouter()
	w := util.PerformRequest(router, "GET", "/")
	assert.Equal(t, http.StatusNotFound, w.Code)
}
