package server

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"vault-proxy/util"
)

func TestHealth(t *testing.T) {
	router := NewRouter()
	w := util.PerformRequest(router, "GET", "/health")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSecrets(t *testing.T) {
	router := NewRouter()
	w := util.PerformRequest(router, "GET", "/secrets")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Contains(t, response["message"], "Success.")
	assert.Contains(t, w.Body.String(), "key1")
}
