package main

import (
	c "baseProject/container"
	"baseProject/server"
	"fmt"
)

func main() {
	container := c.BuildContainer()
	err := container.Invoke(func(server *server.HttpServer) {
		server.Run()
	})
	fmt.Println(err)
	_ = container.Invoke(c.CloseResource)
}
