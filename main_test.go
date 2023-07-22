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

	// TODO: Remove log
	t.Logf(w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resp)

	// assert field of response body
	assert.Nil(t, err)

	// TODO: Assert response
	// assert.Equal(t, float64(id), resp["id"])
	// assert.Equal(t, "John", resp["firstName"])
	// assert.Equal(t, "Doe", resp["lastName"])
	// assert.Equal(t, float64(25), resp["age"])
	// assert.Equal(t, "john.doe@mail.com", resp["email"])
}
