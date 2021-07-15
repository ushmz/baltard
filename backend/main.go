package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ymmt3-lab/koolhaas/backend/database"
	"github.com/ymmt3-lab/koolhaas/backend/handler"
	mw "github.com/ymmt3-lab/koolhaas/backend/middleware"
	"github.com/ymmt3-lab/koolhaas/backend/router"
)

func main() {
	r := router.New()

	d := database.New()
	h := &handler.Handler{DB: d}

	v1 := r.Group("/v1")
	v1.Use(mw.Auth())

	// users
	r.POST("/users", h.CreateUser)
	v1.GET("/users/code/:id", h.GetCompletionCode)

	// task
	v1.GET("/task/:id", h.FetchTaskInfo)
	v1.GET("/serp/:id", h.FetchSerpByID)
	v1.POST("/task/answer", h.SubmitTaskAnswer)

	// logs
	// v1.POST("/users/:userId/logs", h.CreateLog)
	v1.POST("/users/logs/time", h.CreateTaskTimeLog)
	v1.POST("/users/logs/click", h.CreateSerpClickLog)

	defer d.Close()

	r.Logger.Fatal(r.Start(":8080"))
}
