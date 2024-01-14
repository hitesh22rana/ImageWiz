package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hitesh22rana/imagewiz/routes"
	"github.com/hitesh22rana/imagewiz/tasks"
)

const redisAddress = "127.0.0.1:6379"

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${remote_ip} -> ${method} ${uri} ${status} ${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.Setup(e)

	tasks.Init(redisAddress)
	defer tasks.Close()

	go func() {
		if err := e.Start(":8000"); err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := e.Shutdown(context.Background()); err != nil {
		e.Logger.Fatal(err)
	}
}
