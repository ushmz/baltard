package handler

import (
	"baltard/api/dao"
	"baltard/api/service"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Log  *Log
	Serp *Serp
	Task *Task
	User *User
}

func NewHandler(db *sqlx.DB) *Handler {
	log := dao.NewLog(db)
	serp := dao.NewSerp(db)
	task := dao.NewTask(db)
	user := dao.NewUser(db)

	logService := service.NewLogService(log)
	serpService := service.NewSerpService(serp)
	taskService := service.NewTaskService(task)
	userService := service.NewUserService(user, task)

	return &Handler{
		Log:  NewLogHandler(logService),
		Serp: NewSerpHandler(serpService),
		Task: NewTaskHandler(taskService),
		User: NewUserHandler(userService),
	}
}
