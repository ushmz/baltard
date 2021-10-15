package mysql

import (
	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"

	"github.com/jmoiron/sqlx"
)

type LogRepositoryImpl struct {
	DB *sqlx.DB
}

func NewLogRepository(db *sqlx.DB) repo.LogRepository {
	return &LogRepositoryImpl{DB: db}
}

// StoreTaskTimeLog : [Deprecated] Logging task time.
// Key (pair of user_id and task_id) doesn't exist, insert new record.
// Key exists, update `time_on_page` value of requested value.
// This method update task time directly with requested value.
// Therefore, if reloading occur in client side, task time value can be reset unintentionally.
func (l LogRepositoryImpl) StoreTaskTimeLog(p *model.TaskTimeLogParamWithTime) error {
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

// CumulateTaskTimeLog : Logging task time.
// Key (pair of user_id and task_id) doesn't exist, insert new record.
// Key exists, increment `time_on_page` value.
func (l LogRepositoryImpl) CumulateTaskTimeLog(p *model.TaskTimeLogParam) error {
	_, err := l.DB.NamedExec(`
		INSERT INTO
			behavior_logs (
				user_id,
				task_id,
				condition_id
			)
		VALUES (
			:user_id, 
			:task_id, 
			:condition_id
		)
		ON DUPLICATE
			KEY UPDATE
				time_on_page = time_on_page + 1, 
				updated_at = CURRENT_TIMESTAMP
	`, p)
	if err != nil {
		return err
	}
	return nil
}

func (l LogRepositoryImpl) StoreSerpClickLog(p *model.SearchPageClickLogParam) error {
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

func (l LogRepositoryImpl) StoreSearchSeeion(s *model.SearchSession) error {
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
