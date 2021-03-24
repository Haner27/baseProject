package main

import (
	c "baseProject/container"
)

func main() {
	container := c.BuildContainer()
	err := container.Invoke(func() {

	})
}
