package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"ratri/handler"
	db "ratri/infra/mysql"
	mw "ratri/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	// e.Use(middleware.Recover())
	e.Use(mw.Logger())
	e.Use(mw.CacheAdapter())
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

	e.GET("/docs/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	api.POST("/users", h.User.CreateUser)

	v1 := api.Group("/v1")
	// v1.Use(mw.Auth())

	// users
	v1.GET("/users/code/:id", h.User.GetCompletionCode)

	// task
	v1.GET("/task/:id", h.Task.FetchTaskInfo)
	v1.POST("/task/answer", h.Task.SubmitTaskAnswer)
	v1.GET("/serp/:id", h.Serp.FetchSerpByID)
	v1.GET("/serp/:id/icon", h.Serp.FetchSerpWithIconByID)
	v1.GET("/serp/:id/ratio", h.Serp.FetchSerpWithRatioByID)

	// logs
	// v1.POST("/logs/time", h.Log.CreateTaskTimeLog)
	v1.POST("/logs/serp", h.Log.CumulateSerpViewingTime)
	v1.GET("/logs/serp/export", h.Log.ExportSerpViewingTime)
	v1.POST("/logs/pageview", h.Log.CumulatePageViewingTime)
	v1.GET("/logs/pageview/export", h.Log.ExportPageViewingTime)
	v1.POST("/logs/events", h.Log.CreateSerpEventLog)
	v1.GET("/logs/events/export", h.Log.ExportSerpEventLog)
	v1.POST("/logs/session", h.Log.StoreSearchSeeion)
	v1.GET("/logs/session/export", h.Log.ExportSearchSeeion)

	return e
}
