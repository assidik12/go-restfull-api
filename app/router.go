package app

import (
	"github.com/assidik12/go-restfull-api/exception"
	user "github.com/assidik12/go-restfull-api/internal/account/controller"
	category "github.com/assidik12/go-restfull-api/internal/category/controller"
	product "github.com/assidik12/go-restfull-api/internal/product/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(CategoryController category.CategoryController, ProductController product.ProductController, userController user.AccountController) *httprouter.Router {

	router := httprouter.New()

	router.GET("/api/categories", CategoryController.FindAll)
	router.GET("/api/categories/:categoryId", CategoryController.FindById)
	router.POST("/api/categories", CategoryController.Create)
	router.PUT("/api/categories/:categoryId", CategoryController.Update)
	router.DELETE("/api/categories/:categoryId", CategoryController.Delete)

	router.GET("/api/products", ProductController.FindAll)
	router.GET("/api/products/:productId", ProductController.FindById)
	router.POST("/api/products", ProductController.Create)
	router.PUT("/api/products/:productId", ProductController.Update)
	router.DELETE("/api/products/:productId", ProductController.Delete)

	router.POST("/api/auth/login", userController.Login)
	router.POST("/api/auth/register", userController.Register)
	router.PUT("/api/auth/update", userController.Update)

	router.PanicHandler = exception.ErrorHandler

	return router

}
