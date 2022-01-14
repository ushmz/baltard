package mysql

import (
	"database/sql"
	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"

	"github.com/jmoiron/sqlx"
)

type SerpReporitoryImpl struct {
	DB *sqlx.DB
}

func NewSerpRepository(db *sqlx.DB) repo.SerpRepository {
	return &SerpReporitoryImpl{DB: db}
}

func (s SerpReporitoryImpl) FetchSerpByTaskID(taskId, offset int) (*[]model.SearchPage, error) {
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

func (s SerpReporitoryImpl) FetchSerpWithRatioByTaskID(taskId, offset, top int) (*[]model.SerpWithRatioQueryResult, error) {
	swr := []model.SerpWithRatioQueryResult{}
	err := s.DB.Select(&swr, `
		SELECT
			relation_count.id,
			relation_count.title,
			relation_count.url,
			relation_count.snippet,
			relation_count.category,
			ROW_NUMBER() OVER(
				PARTITION BY relation_count.id
				ORDER BY relation_count.category_count
				DESC
			) category_rank,
			relation_count.category_count,
			relation_count.similarweb_count linked_page_count,
			relation_count.category_count / relation_count.similarweb_count category_ratio
		FROM (
			SELECT DISTINCT
				search_pages.id,
				search_pages.title,
				search_pages.url,
				search_pages.snippet,
				similarweb_categories.category,
				COUNT(*) OVER(
					PARTITION BY search_pages.id
				) similarweb_count,
				COUNT(*) OVER(
					PARTITION BY search_pages.id, similarweb_categories.category
				) category_count
			FROM ( SELECT
				page_id,
				similarweb_id,
				idf
			FROM
				search_page_similarweb_relation
			WHERE
				page_id IN (SELECT * FROM (
					SELECT
						page_id
					FROM
						search_page_similarweb_relation
					WHERE
						task_id = ?
					GROUP BY
						page_id
					LIMIT ?, 10
					) as result)
				) as relation
				JOIN
					search_pages
				ON
					relation.page_id = search_pages.id
				JOIN
					similarweb_pages
				ON
					relation.similarweb_id = similarweb_pages.id
				JOIN
					similarweb_categories
				ON
					similarweb_pages.category = similarweb_categories.id
			) as relation_count
	`, taskId, offset*10)
	if err != nil {
		if err == sql.ErrNoRows {
			return &swr, model.NoSuchDataError{}
		}
		return nil, err
	}

	// `sqlx.DB.Select` does not throw `ErrNoRows`,
	// so if length of fetched array is 0, return `model.NoSuchDataError`
	if len(swr) == 0 {
		return &swr, model.NoSuchDataError{}
	}

	return &swr, nil
}
