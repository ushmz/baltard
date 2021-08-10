package handler

import (
	"baltard/api/dao"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Answer    dao.Answer
	Condition dao.Condition
	Log       dao.Log
	Serp      dao.Serp
	Task      dao.Task
	User      dao.User
}

func NewHandler(db *sqlx.DB) *Handler {
	answer := dao.NewAnswer(db)
	condition := dao.NewCondition(db)
	log := dao.NewLog(db)
	serp := dao.NewSerp(db)
	task := dao.NewTask(db)
	user := dao.NewUser(db)

	return &Handler{
		Answer:    answer,
		Condition: condition,
		Log:       log,
		Serp:      serp,
		Task:      task,
		User:      user,
	}
}
