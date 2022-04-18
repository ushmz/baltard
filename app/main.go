package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"ratri/config"
	"ratri/handler"
	fa "ratri/infra/firebase"
	db "ratri/infra/mysql"
	mw "ratri/middleware"

	firebase "firebase.google.com/go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize config values
	if err := config.Init(); err != nil {
		panic(fmt.Sprintf("Failed to get env.yaml : %+v\n", err))
	}

	// Establish DB connection
	d, err := db.New()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to DB: %+v\n", err))
	}
	defer d.Close()

	// Set up firebase SDK
	app, err := fa.InitApp()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize firebase app: %+v\n", err))
	}

	// Create new echo.Echo instance and start listening server
	r := newRouter(d, app)
	conf := config.GetConfig()
	go func() {
		if err := r.Start(":" + conf.GetString("server.port")); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("Failed to start server : %+v\n", err))
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

func newRouter(d *sqlx.DB, app *firebase.App) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	conf := config.GetConfig()
	e.Use(mw.Logger())
	// e.Use(mw.CacheAdapter())
	e.Use(mw.CORSConfig(conf.GetStringSlice("server.cors")))

	h := handler.NewHandler(d, app)

	api := e.Group("/api")
	api.POST("/users", h.User.CreateUser)
	api.POST("/session", h.User.CreateSession)

	v1 := api.Group("/v1")
	// v1.Use(mw.Auth(app))

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
	v1.POST("/logs/serp", h.Log.CumulateSerpDwellTime)
	v1.GET("/logs/serp/export", h.Log.ExportSerpDwellTime)
	v1.POST("/logs/pageview", h.Log.CumulatePageDwellTime)
	v1.GET("/logs/pageview/export", h.Log.ExportPageDwellTime)
	v1.POST("/logs/events", h.Log.CreateSerpEventLog)
	v1.GET("/logs/events/export", h.Log.ExportSerpEventLog)
	v1.POST("/logs/session", h.Log.StoreSearchSeeion)
	v1.GET("/logs/session/export", h.Log.ExportSearchSeeion)

	return e
}
