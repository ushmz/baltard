package mysql

import (
	"fmt"
	"ratri/domain/model"
	repo "ratri/domain/repository"

	"github.com/jmoiron/sqlx"
)

// LogRepositoryImpl : Struct for DB operation
type LogRepositoryImpl struct {
	DB *sqlx.DB
}

// NewLogRepository : Return new struct for DB operation
func NewLogRepository(db *sqlx.DB) repo.LogRepository {
	return &LogRepositoryImpl{DB: db}
}

// FetchAllSerpDwellTimeLogs : Fetch all `SerpDwellLog` data.
// Please make sure that this method is used only for exporting data.
func (l *LogRepositoryImpl) FetchAllSerpDwellTimeLogs() ([]model.SerpDwellTimeLog, error) {
	if l == nil {
		return nil, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

	data := []model.SerpDwellTimeLog{}
	err := l.DB.Select(&data, "SELECT * FROM logs_serp_dwell_time")
	if err != nil {
		return data, fmt.Errorf("Try to get SERP dwell time log: %w", err)
	}
	return data, nil
}

// FetchAllPageDwellTimeLogs : Fetch all `PageDwellLog` data.
// Please make sure that this method is used only for exporting data.
func (l *LogRepositoryImpl) FetchAllPageDwellTimeLogs() ([]model.PageDwellTimeLog, error) {
	if l == nil {
		return nil, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

	data := []model.PageDwellTimeLog{}
	err := l.DB.Select(&data, "SELECT * FROM logs_page_dwell_time")
	if err != nil {
		return data, fmt.Errorf("Try to get result page dwell time log: %w", err)
	}
	return data, nil
}

// FetchAllSerpEventLogs : Fetch all `SerpEventLog` data.
// Please make sure that this method is used only for exporting data.
func (l *LogRepositoryImpl) FetchAllSerpEventLogs() ([]model.SearchPageEventLog, error) {
	if l == nil {
		return nil, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

	data := []model.SearchPageEventLog{}
	err := l.DB.Select(&data, "SELECT * FROM logs_event")
	if err != nil {
		return data, fmt.Errorf("Try to get SERP event logs: %w", err)
	}
	return data, nil
}

// FetchAllSearchSessions : Fetch all `SearchSession` data.
// Please make sure that this method is used only for exporting data.
func (l *LogRepositoryImpl) FetchAllSearchSessions() ([]model.SearchSession, error) {
	if l == nil {
		return nil, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

	data := []model.SearchSession{}
	err := l.DB.Select(&data, "SELECT * FROM search_session")
	if err != nil {
		return data, fmt.Errorf("Try to get search sessions: %w", err)
	}
	return data, nil
}

// CumulateSerpDwellTime : "Upsert" serp viewing time log.
// Key (pair of user_id and task_id) doesn't exist, insert new record.
// Key exists, increment `time_on_page` value.
func (l *LogRepositoryImpl) CumulateSerpDwellTime(p *model.SerpDwellTimeLogParam) error {
	if l == nil {
		return model.ErrNilReceiver
	}

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
	`, &p)
	if err != nil {
		return fmt.Errorf("Try to store SERP dwell time log: %w", err)
	}
	return nil
}

// CumulatePageDwellTime : "Upsert" page viewing time log.
// Key (pair of user_id, task_id and page_id) doesn't exist, insert new record.
// Key exists, increment `time_on_page` value.
func (l *LogRepositoryImpl) CumulatePageDwellTime(p *model.PageDwellTimeLogParam) error {
	if l == nil {
		return fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
	`, &p)
	if err != nil {
		return fmt.Errorf("Try to store result page dwell time log(%v): %w", p, err)
	}
	return nil
}

// StoreSerpEventLog : Insert new SERP event logs
func (l *LogRepositoryImpl) StoreSerpEventLog(p *model.SearchPageEventLogParam) error {
	if l == nil {
		return fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
		)`, &p)
	if err != nil {
		return fmt.Errorf("Try to store SERP event log(%v): %w", p, err)
	}
	return nil
}

// StoreSearchSeeion : Upsert searh session log.
// Insert new row if the user start search session.
// Update "ended_at" field value if the user end search session.
func (l *LogRepositoryImpl) StoreSearchSeeion(s *model.SearchSessionParam) error {
	if l == nil {
		return fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
		`, &s)
	if err != nil {
		return fmt.Errorf("Try to store search session(%v): %w", s, err)
	}
	return nil
}
