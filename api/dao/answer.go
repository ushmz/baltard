package dao

import (
	"baltard/api/models"

	"github.com/jmoiron/sqlx"
)

type Answer interface {
	SubmitTaskAnswer(*models.TaskAnswer) error
}

type AnswerImpl struct {
	DB *sqlx.DB
}

func NewAnswer(db *sqlx.DB) Answer {
	return &AnswerImpl{DB: db}
}

func (a AnswerImpl) SubmitTaskAnswer(answer *models.TaskAnswer) error {
	_, err := a.DB.NamedExec(`
		INSERT INTO
			answers (
				user_id,
				task_id,
				condition_id,
				answer,
				reason
			)
		VALUES (
			:user_id,
			:task_id,
			:condition_id,
			:answer,
			:reason
		)`, answer)
	if err != nil {
		return err
	}
	return nil
}
