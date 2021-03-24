package container

import (
	"baseProject/config"
	"baseProject/datasource"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(config.InitConfig)
	_ = container.Provide(datasource.InitMysqlDB)
	return container
}
