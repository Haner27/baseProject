package server

import (
	"baseProject/config"
	"baseProject/router"
	"baseProject/util/logger"
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
	logger *logger.Logger
	router *router.Router
}

func NewHttpServer(conf *config.Config, logger *logger.Logger, router *router.Router) *HttpServer {
	return &HttpServer{conf, logger, router}
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
		hc.logger.Infof("[%s]server starting on %s", hc.conf.Project.Stage, address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			hc.logger.Errorf("listen: %s", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	hc.logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		hc.logger.Errorf("Server forced to shutdown:%s", err)
	}
}