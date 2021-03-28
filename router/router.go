package router

import (
	"baseProject/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	Engine   *gin.Engine
}

func NewRouter(
	engine *gin.Engine,
	userCtrl *controller.UserController,
) *Router {

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := engine.Group("/v1")
	{
		userGroup := v1.Group("/user")
		{
			userGroup.POST("/register", userCtrl.Register)
			userGroup.POST("/GetUserById", userCtrl.GetUserById)
			userGroup.POST("/GetUsers", userCtrl.GetUsers)
		}
	}

	return &Router{
		engine,
	}
}
