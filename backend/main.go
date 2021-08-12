package main

import (
	"koolhaas/backend/database"
	"koolhaas/backend/handler"
	mw "koolhaas/backend/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover())
	// e.Use(mw.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	d := database.New()
	h := &handler.Handler{DB: d}

	v1 := e.Group("/v1")
	v1.Use(mw.Auth())

	// users
	e.POST("/users", h.CreateUser)
	v1.GET("/users/code/:id", h.GetCompletionCode)

	// task
	v1.GET("/task/:id", h.FetchTaskInfo)
	v1.GET("/serp/:id", h.FetchSerpByID)
	v1.GET("/serp/:id/icon", h.FetchSerpWithIconByID)
	v1.GET("/serp/:id/pct", h.FetchSerpWithDistributionByID)
	v1.POST("/task/answer", h.SubmitTaskAnswer)

	// logs
	// v1.POST("/users/:userId/logs", h.CreateLog)
	v1.POST("/users/logs/time", h.CreateTaskTimeLog)
	v1.POST("/users/logs/click", h.CreateSerpClickLog)
	v1.POST("/task/session", h.StoreSearchSeeion)

	defer d.Close()

	e.Logger.Fatal(e.Start(":8080"))
}
