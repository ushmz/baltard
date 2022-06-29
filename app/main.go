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
	// "github.com/labstack/gommon/log"
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

func httpErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		// Errors not use *echo.HTTPError, such as panic.
		c.Logger().Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	switch he.Message.(type) {
	case handler.ErrWithMessage:
		em := he.Message.(handler.ErrWithMessage)
		c.Logger().Error(em.Error())
		c.JSON(he.Code, em.Why)
	case error:
		e := he.Message.(error)
		c.Logger().Error(e.Error())
		c.NoContent(he.Code)
	default:
		// Unreachable
		c.Logger().Error(he.Message)
		c.JSON(http.StatusInternalServerError, "Unknown error")
	}
}

func newRouter(d *sqlx.DB, app *firebase.App) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = httpErrorHandler
	e.Logger.SetOutput(os.Stderr)

	conf := config.GetConfig()
	if env := conf.GetString("env"); env == "dev" {
		e.Logger.SetHeader("[${level}]${message}")
		e.Logger.SetOutput(os.Stdout)
	}

	// If you store access log with external web server like nginx,
	// You don't need to use this log middleware.
	// e.Use(mw.Logger())

	// If you handle caches with external web server like nginx,
	// You don't need to use this cache middleware.
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
