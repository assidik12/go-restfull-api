//go:build wireinject
// +build wireinject

package config

import (
	"net/http"

	handler "github.com/assidik12/go-restfull-api/internal/delivery/http/handler"
	"github.com/assidik12/go-restfull-api/internal/delivery/http/middleware"
	"github.com/assidik12/go-restfull-api/internal/delivery/http/route"
	"github.com/assidik12/go-restfull-api/internal/infrastructure"
	mysql "github.com/assidik12/go-restfull-api/internal/repository/mysql"
	service "github.com/assidik12/go-restfull-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var validatorSet = wire.NewSet(
	validator.New,
	wire.Value([]validator.Option{}),
)

var userSet = wire.NewSet(
	mysql.NewUserRepository,
	service.NewUserService,
	handler.NewUserHandler,
)

var productSet = wire.NewSet(
	mysql.NewProductRepository,
	service.NewProductService,
	handler.NewProductHandler,
)

var transactionSet = wire.NewSet(
	mysql.NewTransactionRepository,
	service.NewTrancationService,
	handler.NewTransactionHandler,
)

func InitializedServer() *http.Server {
	wire.Build(
		infrastructure.DatabaseConnection,
		validatorSet,
		userSet, productSet, transactionSet,
		route.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
