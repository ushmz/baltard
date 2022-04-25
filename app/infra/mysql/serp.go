package mysql

import (
	"ratri/domain/model"
	repo "ratri/domain/repository"

	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
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
	if s == nil {
		return nil, xerrors.Errorf("SerpReporitoryImpl.FetchSerpByTaskID() is called with nil receiver: %w", model.ErrNoSuchData)
	}

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
		return nil, xerrors.Errorf("Failed to get SERP with taskID(%d), offset(%d): %w", taskID, offset, err)
	}

	// `sqlx.DB.Select` does not throw `ErrNoRows`,
	// so if length of fetched array is 0, return `model.NoSuchDataError`
	if len(srp) == 0 {
		return srp, xerrors.Errorf("SERP with taskID(%d), offset(%d) is not found: %w", taskID, offset, err)
	}

	return srp, nil
}
