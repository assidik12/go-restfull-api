package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductSuccess(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"id":1 ,"name": "Laptop","price": 2000000,"stock": 10,"description": "Laptop Gaming","img": "laptop.png","category_id": 1}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 200, responsee.StatusCode)
}

func TestCreateProductFailed(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":3 ,"name": "Laptop"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 401, responsee.StatusCode)

}

func TestUpdateProductSucces(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":2 ,"name": "Iphone"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/products/2", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 200, responsee.StatusCode)
}

func TestUpdateProductFailed(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":3 ,"name": "Iphone12"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/products/2", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 401, responsee.StatusCode)
}
func TestGetProductSucces(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":3 ,"name": "Iphone12"}`)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/products", requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 200, responsee.StatusCode)
}
func TestGetProductFailed(t *testing.T) {
	router := SetupRouter()
	requestBody := strings.NewReader(`{"id":3 ,"name": "Iphone12"}`)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/products", requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, 401, responsee.StatusCode)
}
