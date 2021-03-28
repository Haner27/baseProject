package main

import (
	c "baseProject/container"
	d "baseProject/datasource"
	m "baseProject/model"
	"baseProject/server"
	"fmt"
)

// @title baseProject api document
// @version 1.0
// @description This is baseProject api document.
// @termsOfService http://swagger.io/terms/

// @contact.name nengfang.han
// @contact.url http://www.swagger.io/support
// @contact.email 369685930@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host baseProject.swagger.io
// @BasePath /v1
func main() {
	container := c.BuildContainer()
	_ = container.Invoke(m.MysqlMigrate)
	err := container.Invoke(func(server *server.HttpServer) {
		server.Run()
	})
	fmt.Println(err)
	_ = container.Invoke(d.CloseResource)
}
