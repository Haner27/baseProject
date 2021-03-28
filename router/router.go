package router

import (
	"baseProject/controller"
	_ "baseProject/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(
	engine *gin.Engine,
	userCtrl *controller.UserController,
) *Router {
	// global middleware
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// swagger api
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ping
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
