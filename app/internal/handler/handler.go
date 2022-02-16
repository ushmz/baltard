package handler

import (
	"ratri/internal/infra/fileio"
	db "ratri/internal/infra/mysql"
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
	linkedPage := db.NewLinkedPageRepository(dbx)
	log := db.NewLogRepository(dbx)
	serp := db.NewSerpRepository(dbx)
	task := db.NewTaskRepository(dbx)
	user := db.NewUserRepository(dbx)
	file := fileio.NewLogStore()

	logService := usecase.NewLogUsecase(log, file)
	serpService := usecase.NewSerpUsecase(serp, linkedPage)
	taskService := usecase.NewTaskUsecase(task)
	userService := usecase.NewUserUsecase(user, task)

	return &Handler{
		Log:  NewLogHandler(logService),
		Serp: NewSerpHandler(serpService),
		Task: NewTaskHandler(taskService),
		User: NewUserHandler(userService),
	}
}
