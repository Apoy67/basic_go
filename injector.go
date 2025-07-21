//go:build wireinject
// +build wireinject

package main

import (
	"basic-restfull-golang/app"
	"basic-restfull-golang/controller"
	"basic-restfull-golang/middleware"
	"basic-restfull-golang/repository"
	"basic-restfull-golang/service"
	"basic-restfull-golang/test"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

var validatorSet = wire.NewSet(test.NewValidator)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validatorSet,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
