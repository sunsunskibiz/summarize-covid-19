package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCovidSummary(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/covid/summary", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resp)

	assert.Nil(t, err)
	assert.NotEmpty(t, resp["Province"])
	assert.NotEmpty(t, resp["Province"].(map[string]interface{})["Bangkok"])
	assert.NotEmpty(t, resp["AgeGroup"])
	assert.NotEmpty(t, resp["AgeGroup"].(map[string]interface{})["0-30"])
}
