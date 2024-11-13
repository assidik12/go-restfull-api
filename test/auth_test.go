package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthRegisterSucces(t *testing.T) {
	router := SetupTestRouter()

	requestBody := strings.NewReader(`{"username": "sidiktest","email": "admin@gamil.com", "password": "admin12345"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/auth/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, http.StatusCreated, responsee.StatusCode)
}
func TestAuthRegisterFailed(t *testing.T) {
	router := SetupTestRouter()

	requestBody := strings.NewReader(`{ "password": "admin"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/auth/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, http.StatusBadRequest, responsee.StatusCode)
}

func TestAuthLoginSucces(t *testing.T) {
	router := SetupTestRouter()

	requestBody := strings.NewReader(`{"email": "admin@gamil.com", "password": "admin12345"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/auth/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, http.StatusOK, responsee.StatusCode)
}

func TestAuthLoginFailed(t *testing.T) {
	router := SetupTestRouter()

	requestBody := strings.NewReader(`{"email": "admin@gamil.com", "password": "salah"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/auth/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, http.StatusUnauthorized, responsee.StatusCode)
}

func TestAuthUpdateSucces(t *testing.T) {
	router := SetupTestRouter()

	requestBody := strings.NewReader(`{"email": "admin@gamil.com", "password": "admin12345"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/auth/update", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusAccepted, response.StatusCode)

}

func TestAuthUpdateFailed(t *testing.T) {
	router := SetupTestRouter()

	requestBody := strings.NewReader(`{"email": "salah@gamil.com", "password": "salah"}`)

	request := httptest.NewRequest(http.MethodPut, "ttp://localhost:3000/api/auth/update", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}
