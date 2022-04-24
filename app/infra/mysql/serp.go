package mysql

import (
	"database/sql"
	"ratri/domain/model"
	repo "ratri/domain/repository"

	"github.com/jmoiron/sqlx"
)

// SerpReporitoryImpl : Implemention of SERP repository
type SerpReporitoryImpl struct {
	DB *sqlx.DB
}

// NewSerpRepository : Return new SERP repository struct
func NewSerpRepository(db *sqlx.DB) repo.SerpRepository {
	return &SerpReporitoryImpl{DB: db}
}

// FetchSerpByTaskID : Get result pages by task ID
func (s *SerpReporitoryImpl) FetchSerpByTaskID(taskID, offset int) ([]model.SearchPage, error) {
	srp := []model.SearchPage{}
	// [TODO] Performance measure.
	// srp := make([]model.SearchPage, 10)
	err := s.DB.Select(&srp, `
		SELECT
			search_pages.id,
			search_pages.title,
			search_pages.url,
			search_pages.snippet
		FROM
			search_pages
		WHERE
			task_id = ?
		LIMIT
			?, 10
	`, taskID, offset*10)
	if err != nil {
		if err == sql.ErrNoRows {
			return srp, model.ErrNoSuchData
		}
		return nil, err
	}

	// `sqlx.DB.Select` does not throw `ErrNoRows`,
	// so if length of fetched array is 0, return `model.NoSuchDataError`
	if len(srp) == 0 {
		return srp, model.ErrNoSuchData
	}

	return srp, nil
}
