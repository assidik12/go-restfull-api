package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCategorySuccess(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"id":1 ,"name": "Laptop"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 200, responsee.StatusCode)
}

func TestCreateCategoryFailed(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":3 ,"name": "Laptop"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 401, responsee.StatusCode)

}

func TestUpdateCategorySucces(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":2 ,"name": "Iphone"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/2", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 200, responsee.StatusCode)
}

func TestUpdateCategoryFailed(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":3 ,"name": "Iphone12"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/2", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 401, responsee.StatusCode)
}
func TestGetCategorySucces(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":3 ,"name": "Iphone12"}`)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 200, responsee.StatusCode)
}
func TestGetCategoryFailed(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":3 ,"name": "Iphone12"}`)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 401, responsee.StatusCode)
}
