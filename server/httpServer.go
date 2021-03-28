package server

import (
	"baseProject/config"
	"baseProject/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HttpServer struct {
	conf *config.Config
	router *router.Router
}

func NewHttpServer(conf *config.Config, router *router.Router) *HttpServer {
	return &HttpServer{conf, router}
}

func (hc *HttpServer) GetAddress() string {
	return fmt.Sprintf("0.0.0.0:%d", hc.conf.Project.Port)
}

func (hc *HttpServer) Run() {
	if hc.conf.Project.Stage == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	address := hc.GetAddress()
	srv := &http.Server{
		Addr: address,
		Handler: hc.router.Engine,
	}
	go func() {
		fmt.Printf("[%s]server starting on %s", hc.conf.Project.Stage, address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown:%s", err)
	}
}