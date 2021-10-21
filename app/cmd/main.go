package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"ratri/internal/handler"
	db "ratri/internal/infra/mysql"

	mw "ratri/internal/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	d := db.New()
	defer d.Close()
	r := NewRouter(d)
	go func() {
		if err := r.Start(":8080"); err != nil && err != http.ErrServerClosed {
			r.Logger.Fatal("Shutting down the server")
		}
	}()

	// Gracefully shutdown the server with a timeout (10 seconds.)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Shutdown(ctx); err != nil {
		r.Logger.Fatal(err)
	}
}

func NewRouter(d *sqlx.DB) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover())
	e.Use(mw.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowMethods: []string{
			echo.GET,
			echo.HEAD,
			echo.PUT,
			echo.PATCH,
			echo.POST,
			echo.DELETE,
		},
	}))

	h := handler.NewHandler(d)

	v1 := e.Group("/v1")
	// v1.Use(mw.Auth())

	// users
	e.POST("/users", h.User.CreateUser)
	v1.GET("/users/code/:id", h.User.GetCompletionCode)

	// task
	v1.GET("/task/:id", h.Task.FetchTaskInfo)
	v1.GET("/serp/:id", h.Serp.FetchSerpByID)
	v1.GET("/serp/:id/icon", h.Serp.FetchSerpWithIconByID)
	v1.GET("/serp/:id/pct", h.Serp.FetchSerpWithDistributionByID)
	v1.POST("/task/answer", h.Task.SubmitTaskAnswer)

	// logs
	v1.PATCH("/users/logs/time", h.Log.CreateTaskTimeLog)
	v1.POST("/users/logs/time", h.Log.CumulateTaskTimeLog)
	v1.POST("/users/logs/click", h.Log.CreateSerpClickLog)
	v1.POST("/task/session", h.Log.StoreSearchSeeion)

	return e
}
