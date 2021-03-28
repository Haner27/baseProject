package main

import (
	c "baseProject/container"
	d "baseProject/datasource"
	m "baseProject/model"
	"baseProject/server"
	"fmt"
)

func main() {
	container := c.BuildContainer()
	_ = container.Invoke(m.MysqlMigrate)
	err := container.Invoke(func(server *server.HttpServer) {
		server.Run()
	})
	fmt.Println(err)
	_ = container.Invoke(d.CloseResource)
}
