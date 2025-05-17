//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"renaksiopdService/app"
	"renaksiopdService/controller"
	"renaksiopdService/middleware"
	"renaksiopdService/repository"
	"renaksiopdService/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var rencanaAksiOpdSet = wire.NewSet(
	repository.NewRencanaAksiOpdRepositoryImpl,
	wire.Bind(new(repository.RencanaAksiOpdRepository), new(*repository.RencanaAksiOpdRepositoryImpl)),
	service.NewRencanaAksiOpdServiceImpl,
	wire.Bind(new(service.RencanaAksiOpdService), new(*service.RencanaAksiOpdServiceImpl)),
	controller.NewRencanaAksiOpdControllerImpl,
	wire.Bind(new(controller.RencanaAksiOpdController), new(*controller.RencanaAksiOpdControllerImpl)),
)

func InitializeServer() *http.Server {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		rencanaAksiOpdSet,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		app.NewRouter,
		NewServer,
	)
	return nil
}
