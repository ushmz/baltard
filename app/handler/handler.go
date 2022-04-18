package handler

import (
	"ratri/infra/fileio"
	fa "ratri/infra/firebase"
	db "ratri/infra/mysql"
	"ratri/usecase"

	firebase "firebase.google.com/go"
	"github.com/jmoiron/sqlx"
)

// Handler : Handle HTTP requests
type Handler struct {
	Log  *Log
	Serp *Serp
	Task *Task
	User *User
}

// NewHandler : Return new handler struct
func NewHandler(dbx *sqlx.DB, app *firebase.App) *Handler {
	linkedPage := db.NewLinkedPageRepository(dbx)
	log := db.NewLogRepository(dbx)
	serp := db.NewSerpRepository(dbx)
	task := db.NewTaskRepository(dbx)
	user := db.NewUserRepository(dbx)
	file := fileio.NewLogStore()
	userAuth := fa.NewUserAuthenticationImpl(app)

	logService := usecase.NewLogUsecase(log, file)
	serpService := usecase.NewSerpUsecase(serp, linkedPage)
	taskService := usecase.NewTaskUsecase(task)
	userService := usecase.NewUserUsecase(user, task, userAuth)

	return &Handler{
		Log:  NewLogHandler(logService),
		Serp: NewSerpHandler(serpService),
		Task: NewTaskHandler(taskService),
		User: NewUserHandler(userService),
	}
}
