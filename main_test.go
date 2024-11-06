package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Error struct {
	Success bool
	Message string
}

func TestConvertRouteInvalidRequest(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/convert", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestConvertRouteMissingToCurrency(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/convert?from=USD", nil)
	router.ServeHTTP(w, req)

	resBody, _ := io.ReadAll(w.Body)
	errorObj := &Error{}

	json.Unmarshal(resBody, &errorObj)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, false, errorObj.Success)
	assert.Equal(t, "'to' currency not included in request", errorObj.Message)
}

func TestConvertRouteInvalidToCurrency(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/convert?from=USD&to=XYZ", nil)
	router.ServeHTTP(w, req)

	resBody, _ := io.ReadAll(w.Body)
	errorObj := &Error{}

	json.Unmarshal(resBody, &errorObj)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, false, errorObj.Success)
	assert.Equal(t, "'to' currency not a supported currency", errorObj.Message)
}
