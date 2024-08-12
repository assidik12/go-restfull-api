//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/assidik12/go-restfull-api/app"
	categoryCtrl "github.com/assidik12/go-restfull-api/internal/category/controller"
	categoryRepo "github.com/assidik12/go-restfull-api/internal/category/repository"
	categorySrvc "github.com/assidik12/go-restfull-api/internal/category/service"
	productCtrl "github.com/assidik12/go-restfull-api/internal/product/controller"
	productRepo "github.com/assidik12/go-restfull-api/internal/product/repository"
	productServ "github.com/assidik12/go-restfull-api/internal/product/service"
	"github.com/assidik12/go-restfull-api/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	categoryRepo.NewCategoryRepository,
	wire.Bind(new(categoryRepo.CategoryRepository), new(*categoryRepo.CategoryRepositoryImpl)),
	categorySrvc.NewCategoryService,
	wire.Bind(new(categorySrvc.CategoryService), new(*categorySrvc.CategoryServiceImpl)),
	categoryCtrl.NewCategoryController,
	wire.Bind(new(categoryCtrl.CategoryController), new(*categoryCtrl.CategoryControllerImpl)),
	wire.Value([]validator.Option{}),
)

var productSet = wire.NewSet(
	productRepo.NewProductRepository,
	wire.Bind(new(productRepo.ProductRepository), new(*productRepo.ProductRepositoryImpl)),
	productServ.NewProductService,
	wire.Bind(new(productServ.ProductService), new(*productServ.ProductServiceImpl)),
	productCtrl.NewProductController,
	wire.Bind(new(productCtrl.ProductController), new(*productCtrl.ProductControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet, productSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
