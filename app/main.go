package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	d := db.New()
	defer d.Close()

	// Set up firebase SDK
	app, err := fa.InitApp()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize firebase app: %+v\n", err))
	}

	r := NewRouter(d, app)
	go func() {
		if err := r.Start(":8080"); err != nil && err != http.ErrServerClosed {
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

func NewRouter(d *sqlx.DB, app *firebase.App) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(mw.Logger())
	e.Use(mw.CacheAdapter())
	e.Use(mw.CORSConfig())

	h := handler.NewHandler(d, app)

	api := e.Group("/api")
	api.POST("/users", h.User.CreateUser)
	api.POST("/session", h.User.CreateSession)

	v1 := api.Group("/v1")
	v1.Use(mw.Auth(app))

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
