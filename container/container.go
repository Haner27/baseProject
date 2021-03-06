package container

import (
	"baseProject/config"
	"baseProject/controller"
	"baseProject/datasource"
	"baseProject/middleware"
	"baseProject/repository"
	"baseProject/router"
	"baseProject/server"
	"baseProject/service"
	"baseProject/util/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	// config
	_ = container.Provide(config.InitConfig)

	// logger
	_ = container.Provide(logger.NewLogger)

	// datasource
	_ = container.Provide(datasource.NewMysqlDB)
	_ = container.Provide(datasource.NewRedisDb)

	// repository
	_ = container.Provide(repository.NewUserRepository)

	// service
	_ = container.Provide(service.NewUserService)

	// controller
	_ = container.Provide(controller.NewUserController)

	// gin engine
	_ = container.Provide(gin.New)

	// middleware
	_ = container.Provide(middleware.NewSliderWindowLimiter)
	_ = container.Provide(middleware.NewTokenBucketLimiter)

	// router
	_ = container.Provide(router.NewRouter)

	// server
	_ = container.Provide(server.NewHttpServer)
	return container
}
