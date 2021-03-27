package container

import (
	"baseProject/config"
	"baseProject/controller"
	"baseProject/datasource"
	"baseProject/repository"
	"baseProject/router"
	"baseProject/server"
	"baseProject/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	// config
	_ = container.Provide(config.InitConfig)

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

	// router
	_ = container.Provide(router.NewRouter)

	// server
	_ = container.Provide(server.NewHttpServer)
	return container
}
