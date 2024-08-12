//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil

}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabaseMysql, NewDatabasePostgresql, NewDatabaseRepository)
	return nil
}

var fooset = wire.NewSet(NewFooRepository, NewFooService)
var barset = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooset, barset, NweFooBarService)
	return nil
}
