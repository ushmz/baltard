package handler

import (
	"baltard/internal/infra/db"
	"ratri/internal/usecase"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Log  *Log
	Serp *Serp
	Task *Task
	User *User
}

func NewHandler(dbx *sqlx.DB) *Handler {
	log := db.NewLogRepository(dbx)
	serp := db.NewSerpRepository(dbx)
	task := db.NewTaskRepository(dbx)
	user := db.NewUserRepository(dbx)

	logService := usecase.NewLogUsecase(log)
	serpService := usecase.NewSerpUsecase(serp)
	taskService := usecase.NewTaskUsecase(task)
	userService := usecase.NewUserUsecase(user, task)

	return &Handler{
		Log:  NewLogHandler(logService),
		Serp: NewSerpHandler(serpService),
		Task: NewTaskHandler(taskService),
		User: NewUserHandler(userService),
	}
}
