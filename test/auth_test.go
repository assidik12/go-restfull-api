package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthRegisterSucces(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"username": "admin","email": "admin@gamil.com", "password": "admin"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 200, responsee.StatusCode)
}
func TestAuthRegisterFailed(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{ "password": "admin"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 400, responsee.StatusCode)
}

func TestAuthLoginSucces(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"email": "admin@gamil.com", "password": "admin"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 200, responsee.StatusCode)
}

func TestAuthLoginFailed(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"email": "admin@gamil.com", "password": "salah"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 401, responsee.StatusCode)
}

func TestAuthUpdateSucces(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"email": "admin@gamil.com", "password": "admin12345"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/auth/update", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

}

func TestAuthUpdateFailed(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"email": "salah@gamil.com", "password": "salah"}`)

	request := httptest.NewRequest(http.MethodPut, "ttp://localhost:3000/api/auth/update", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}
