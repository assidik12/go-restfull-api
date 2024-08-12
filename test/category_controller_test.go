package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	categoryCtrl "github.com/assidik12/go-restfull-api/internal/category/controller"
	categoryRepo "github.com/assidik12/go-restfull-api/internal/category/repository"
	categoryServ "github.com/assidik12/go-restfull-api/internal/category/service"
	productCtrl "github.com/assidik12/go-restfull-api/internal/product/controller"
	productRepo "github.com/assidik12/go-restfull-api/internal/product/repository"
	productServ "github.com/assidik12/go-restfull-api/internal/product/service"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	"github.com/assidik12/go-restfull-api/app"
	"github.com/assidik12/go-restfull-api/middleware"
)

func SetupRouter() http.Handler {
	db := app.SetupTestDB()

	validate := validator.New()

	categoryRepository := categoryRepo.NewCategoryRepository()
	categoryService := categoryServ.NewCategoryService(categoryRepository, db, validate)
	CategoryController := categoryCtrl.NewCategoryController(categoryService)

	productRepository := productRepo.NewProductRepository()
	productService := productServ.NewProductService(productRepository, db, validate)
	ProductController := productCtrl.NewProductController(productService)

	router := app.NewRouter(CategoryController, ProductController)
	return middleware.NewAuthMiddleware(router)
}

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
