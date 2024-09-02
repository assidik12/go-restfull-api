//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/assidik12/go-restfull-api/app"
	accountCtrl "github.com/assidik12/go-restfull-api/internal/account/controller"
	accountRepo "github.com/assidik12/go-restfull-api/internal/account/repository"
	accountServ "github.com/assidik12/go-restfull-api/internal/account/service"
	categoryCtrl "github.com/assidik12/go-restfull-api/internal/category/controller"
	categoryRepo "github.com/assidik12/go-restfull-api/internal/category/repository"
	categorySrvc "github.com/assidik12/go-restfull-api/internal/category/service"
	productCtrl "github.com/assidik12/go-restfull-api/internal/product/controller"
	productRepo "github.com/assidik12/go-restfull-api/internal/product/repository"
	productServ "github.com/assidik12/go-restfull-api/internal/product/service"
	transactionCtrl "github.com/assidik12/go-restfull-api/internal/transaction/controller"
	transactionRepo "github.com/assidik12/go-restfull-api/internal/transaction/repository"
	transactionServ "github.com/assidik12/go-restfull-api/internal/transaction/service"
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

var accountSet = wire.NewSet(
	accountRepo.NewAccountRepository,
	wire.Bind(new(accountRepo.AccountRepository), new(*accountRepo.AccountRepositoryImpl)),
	accountServ.NewAccountService,
	wire.Bind(new(accountServ.AccountService), new(*accountServ.AccountServiceImpl)),
	accountCtrl.NewAccountController,
	wire.Bind(new(accountCtrl.AccountController), new(*accountCtrl.AccountControllerImpl)),
)

var transactionSet = wire.NewSet(
	transactionRepo.NewTransactionRepository,
	wire.Bind(new(transactionRepo.TransactionRepository), new(*transactionRepo.TransactionRepositoryImpl)),
	transactionServ.NewTransactionService,
	wire.Bind(new(transactionServ.TransactionService), new(*transactionServ.TransactionServiceImpl)),
	transactionCtrl.NewTransactionController,
	wire.Bind(new(transactionCtrl.TransactionController), new(*transactionCtrl.TransactionControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet, productSet, accountSet, transactionSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
