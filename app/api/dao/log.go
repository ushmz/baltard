package dao

import (
	"baltard/api/model"

	"github.com/jmoiron/sqlx"
)

type Log interface {
	CreateTaskTimeLog(*model.TaskTimeLogParam) error
	CreateSerpClickLog(*model.SearchPageClickLogParam) error
	StoreSearchSeeion(*model.SearchSession) error
}

type LogImpl struct {
	DB *sqlx.DB
}

func NewLog(db *sqlx.DB) Log {
	return &LogImpl{DB: db}
}

func (l LogImpl) CreateTaskTimeLog(p *model.TaskTimeLogParam) error {
	_, err := l.DB.NamedExec(`
		INSERT INTO
			behavior_logs (
				user_id,
				task_id,
				time_on_page,
				condition_id
			)
		VALUES (
			:user_id, 
			:task_id, 
			:time_on_page, 
			:condition_id
		)
		ON DUPLICATE
			KEY UPDATE
				time_on_page = :time_on_page, 
				updated_at = CURRENT_TIMESTAMP
	`, p)
	if err != nil {
		return err
	}
	return nil
}

func (l LogImpl) CreateSerpClickLog(p *model.SearchPageClickLogParam) error {
	_, err := l.DB.NamedExec(`
		INSERT INTO
			behavior_logs_click (
				user_id,
				task_id,
				condition_id,
				time_on_page,
				serp_page,
				serp_rank,
				is_visible
			)
		VALUES (
			:user_id, 
			:task_id, 
			:condition_id,
			:time_on_page, 
			:serp_page,
			:serp_rank,
			:is_visible
		)`, p)
	if err != nil {
		return err
	}
	return nil
}

func (l LogImpl) StoreSearchSeeion(s *model.SearchSession) error {
	_, err := l.DB.NamedExec(`
		INSERT INTO
			search_session(
				user_id,
				task_id,
				condition_id
			)
		VALUES (
			:user_id,
			:task_id,
			:condition_id
		)
		ON DUPLICATE KEY
			UPDATE
				ended_at = CURRENT_TIMESTAMP
		`, s)
	if err != nil {
		return err
	}
	return nil
}
