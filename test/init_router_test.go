package test

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/assidik12/go-restfull-api/app"
	accountCtrl "github.com/assidik12/go-restfull-api/internal/account/controller"
	accountRepo "github.com/assidik12/go-restfull-api/internal/account/repository"
	accountServ "github.com/assidik12/go-restfull-api/internal/account/service"
	categoryCtrl "github.com/assidik12/go-restfull-api/internal/category/controller"
	categoryRepo "github.com/assidik12/go-restfull-api/internal/category/repository"
	categoryServ "github.com/assidik12/go-restfull-api/internal/category/service"
	productCtrl "github.com/assidik12/go-restfull-api/internal/product/controller"
	productRepo "github.com/assidik12/go-restfull-api/internal/product/repository"
	productServ "github.com/assidik12/go-restfull-api/internal/product/service"
	transactionCtrl "github.com/assidik12/go-restfull-api/internal/transaction/controller"
	transactionRepo "github.com/assidik12/go-restfull-api/internal/transaction/repository"
	transactionServ "github.com/assidik12/go-restfull-api/internal/transaction/service"

	"github.com/assidik12/go-restfull-api/middleware"
	"github.com/go-playground/validator/v10"
)

func SetupTestRouter() http.Handler {
	db := app.SetupTestDB()

	validate := validator.New()

	categoryRepository := categoryRepo.NewCategoryRepository()
	categoryService := categoryServ.NewCategoryService(categoryRepository, db, validate)
	CategoryController := categoryCtrl.NewCategoryController(categoryService)

	productRepository := productRepo.NewProductRepository()
	productService := productServ.NewProductService(productRepository, db, validate)
	ProductController := productCtrl.NewProductController(productService)

	accountRepository := accountRepo.NewAccountRepository()
	accountService := accountServ.NewAccountService(accountRepository, db, validate)
	accountController := accountCtrl.NewAccountController(accountService)

	transactionRepository := transactionRepo.NewTransactionRepository()
	transactionService := transactionServ.NewTransactionService(transactionRepository, db, validate)
	TransactionController := transactionCtrl.NewTransactionController(transactionService)

	router := app.NewRouter(CategoryController, ProductController, accountController, TransactionController)

	return middleware.NewAuthMiddleware(router)
}
