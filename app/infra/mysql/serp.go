package mysql

import (
	"database/sql"
	"ratri/domain/model"
	repo "ratri/domain/repository"

	"github.com/jmoiron/sqlx"
)

type SerpReporitoryImpl struct {
	DB *sqlx.DB
}

func NewSerpRepository(db *sqlx.DB) repo.SerpRepository {
	return &SerpReporitoryImpl{DB: db}
}

func (s *SerpReporitoryImpl) FetchSerpByTaskID(taskId, offset int) (*[]model.SearchPage, error) {
	srp := []model.SearchPage{}
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
	`, taskId, offset*10)
	if err != nil {
		if err == sql.ErrNoRows {
			return &srp, model.NoSuchDataError{}
		}
		return nil, err
	}

	// `sqlx.DB.Select` does not throw `ErrNoRows`,
	// so if length of fetched array is 0, return `model.NoSuchDataError`
	if len(srp) == 0 {
		return &srp, model.NoSuchDataError{}
	}

	return &srp, nil
}
