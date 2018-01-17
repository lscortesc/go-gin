package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/lscortesc/go-gin/models"
)

func TestHomeRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	var res models.APIResponse
	body := []byte(w.Body.String())
	json.Unmarshal(body, &res)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Data recover successfully", res.Message)
	assert.Equal(t, 200, res.Status)
}
