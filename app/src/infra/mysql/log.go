package mysql

import (
	"ratri/src/domain/model"
	repo "ratri/src/domain/repository"

	"github.com/jmoiron/sqlx"
)

type LogRepositoryImpl struct {
	DB *sqlx.DB
}

func NewLogRepository(db *sqlx.DB) repo.LogRepository {
	return &LogRepositoryImpl{DB: db}
}

// FetchAllSerpViewingTimeLogs : Fetch all `SerpViewingLog` data.
// Please make sure that this method is used only for exporting data.
func (l *LogRepositoryImpl) FetchAllSerpViewingTimeLogs() ([]model.SerpViewingLog, error) {
	data := []model.SerpViewingLog{}
	err := l.DB.Select(&data, "SELECT * FROM logs_serp_dwell_time")
	if err != nil {
		return data, err
	}
	return data, nil
}

// FetchAllPageViewingTimeLogs : Fetch all `PageViewingLog` data.
// Please make sure that this method is used only for exporting data.
func (l *LogRepositoryImpl) FetchAllPageViewingTimeLogs() ([]model.PageViewingLog, error) {
	data := []model.PageViewingLog{}
	err := l.DB.Select(&data, "SELECT * FROM logs_page_dwell_time")
	if err != nil {
		return data, err
	}
	return data, nil
}

// FetchAllSerpEventLogs : Fetch all `SerpEventLog` data.
// Please make sure that this method is used only for exporting data.
func (l *LogRepositoryImpl) FetchAllSerpEventLogs() ([]model.SearchPageEventLog, error) {
	data := []model.SearchPageEventLog{}
	err := l.DB.Select(&data, "SELECT * FROM logs_event")
	if err != nil {
		return data, err
	}
	return data, nil
}

// FetchAllSearchSessions : Fetch all `SearchSession` data.
// Please make sure that this method is used only for exporting data.
func (l *LogRepositoryImpl) FetchAllSearchSessions() ([]model.SearchSession, error) {
	data := []model.SearchSession{}
	err := l.DB.Select(&data, "SELECT * FROM search_session")
	if err != nil {
		return data, err
	}
	return data, nil
}

// CumulateSerpViewingTime : "Upsert" serp viewing time log.
// Key (pair of user_id and task_id) doesn't exist, insert new record.
// Key exists, increment `time_on_page` value.
func (l *LogRepositoryImpl) CumulateSerpViewingTime(p model.SerpViewingLogParam) error {
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

// CumulatePageViewingTime : "Upsert" page viewing time log.
// Key (pair of user_id, task_id and page_id) doesn't exist, insert new record.
// Key exists, increment `time_on_page` value.
func (l *LogRepositoryImpl) CumulatePageViewingTime(p model.PageViewingLogParam) error {
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

func (l *LogRepositoryImpl) StoreSerpEventLog(p model.SearchPageEventLogParam) error {
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

// StoreSearchSeeion : Upsert searh session log.
// Insert new row if the user start search session.
// Update "ended_at" field value if the user end search session.
func (l *LogRepositoryImpl) StoreSearchSeeion(s model.SearchSessionParam) error {
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
