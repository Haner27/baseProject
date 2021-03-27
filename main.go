package main

import (
	c "baseProject/container"
	"baseProject/server"
)

func main() {
	container := c.BuildContainer()
	err := container.Invoke(func(server *server.HttpServer) {
		server.Run()
	})
	err := container.Invoke(func(server *server.HttpServer) {
		server.Run()
	})
}
