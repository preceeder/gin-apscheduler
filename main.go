//   File Name:  main.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/12 17:43
//    Change Activity:

package main

import (
	"context"
	"gin-apscheduler/Init"
	_ "gin-apscheduler/docs"
	"gin-apscheduler/server"
	_ "gin-apscheduler/task"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := server.Router()
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerfiles.Handler, "disableSwagger"))

	Init.Init()
	Run(r, "0.0.0.0:8080", func() {
		Init.Scheduler.Stop()
	})
}

func Run(r *gin.Engine, addr string, stop func()) {

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	//保证下面的优雅启停
	go func() {
		slog.Info("server running ", "addr", "http://"+srv.Addr, "swag", "http://"+srv.Addr+"/swagger/index.html")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
		} else {
			slog.Info("启动uri", "uri", addr)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting Down project ...", "server-name", addr)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	if stop != nil {
		stop()
	}
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("stop error ", "err", err.Error())
	}

	select {
	case <-ctx.Done():
		slog.Info("stop success ")
	}

}
