package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/ymmt3-lab/koolhaas/backend/api/dao"
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
