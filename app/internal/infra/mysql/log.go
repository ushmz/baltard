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

// CumulateSerpViewingTime : Logging task time.
// Key (pair of user_id and task_id) doesn't exist, insert new record.
// Key exists, increment `time_on_page` value.
func (l LogRepositoryImpl) CumulateSerpViewingTime(p *model.SerpViewingLogParam) error {
	_, err := l.DB.NamedExec(`
		INSERT INTO
			logs_serp_dwell_time (
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

// CumulatePageViewingTime : Logging page viewing time.
// Key (pair of user_id, task_id and page_id) doesn't exist, insert new record.
// Key exists, increment `time_on_page` value.
func (l LogRepositoryImpl) CumulatePageViewingTime(p *model.PageViewingLogParam) error {
	_, err := l.DB.NamedExec(`
		INSERT INTO
			logs_page_dwell_time (
				user_id,
				task_id,
				condition_id,
				page_id
			)
		VALUES (
			:user_id, 
			:task_id, 
			:condition_id,
			:page_id
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

func (l LogRepositoryImpl) StoreSerpEventLog(p *model.SearchPageEventLogParam) error {
	_, err := l.DB.NamedExec(`
		INSERT INTO
			logs_event (
				user_id,
				task_id,
				condition_id,
				time_on_page,
				serp_page,
				serp_rank,
				is_visible,
				event
			)
		VALUES (
			:user_id, 
			:task_id, 
			:condition_id,
			:time_on_page, 
			:serp_page,
			:serp_rank,
			:is_visible,
			:event
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
