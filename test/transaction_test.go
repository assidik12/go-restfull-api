package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	router := SetupTestRouter()

	requestBody := strings.NewReader(`{}`)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/transactions/2", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, http.StatusOK, responsee.StatusCode)
}

func TestGetTransactionsFailed(t *testing.T) {
	router := SetupTestRouter()

	requestBody := strings.NewReader(`{}`)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/transactions/5", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, http.StatusNotFound, responsee.StatusCode)
}

func TestCreateTransaction(t *testing.T) {
	router := SetupTestRouter()
	requwstBody := strings.NewReader(`{"total_price":50000, "user_id": 2,"id_trx":12344, "products":[{"id":1, "qty":5},{"id":2, "qty":3}]}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/transactions", requwstBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, http.StatusCreated, responsee.StatusCode)
}
func TestCreateTransactionFailed(t *testing.T) {
	router := SetupTestRouter()
	requwstBody := strings.NewReader(`{"total_price":50000, "user_id": 2,"id_trx":1234214}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/transactions", requwstBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA DONG BRO")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	responsee := recorder.Result()

	assert.Equal(t, http.StatusBadGateway, responsee.StatusCode)
}
