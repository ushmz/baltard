package db

import (
	"baltard/internal/domain/model"
	repo "baltard/internal/domain/repository"

	"github.com/jmoiron/sqlx"
)

type LogRepositoryImpl struct {
	DB *sqlx.DB
}

func NewLogRepository(db *sqlx.DB) repo.LogRepository {
	return &LogRepositoryImpl{DB: db}
}

func (l LogRepositoryImpl) StoreTaskTimeLog(p *model.TaskTimeLogParam) error {
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
