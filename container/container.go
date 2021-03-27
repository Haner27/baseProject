package container

import (
	"baseProject/config"
	"baseProject/controller"
	"baseProject/datasource"
	"baseProject/repository"
	"baseProject/service"
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

	// router

	// server
	return container
}
