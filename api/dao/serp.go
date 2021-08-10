package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/ymmt3-lab/koolhaas/backend/api/models"
)

type Serp interface {
	FetchSerpByID(taskId, offset int) ([]models.SearchPage, error)
	FetchSerpWithIconByID(taskId, offset, top int) ([]models.SerpWithIconQueryResult, error)
	FetchSerpWithDistributionByID(taskId, offset, top int) ([]models.SerpWithDistributionQueryResult, error)
}

type SerpImpl struct {
	DB *sqlx.DB
}

func NewSerp(db *sqlx.DB) Serp {
	return &SerpImpl{DB: db}
}

func (s SerpImpl) FetchSerpByID(taskId, offset int) ([]models.SearchPage, error) {
	srp := []models.SearchPage{}
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
		return nil, err
	}
	return srp, nil
}

func (s SerpImpl) FetchSerpWithIconByID(taskId, offset, top int) ([]models.SerpWithIconQueryResult, error) {
	swi := []models.SerpWithIconQueryResult{}
	err := s.DB.Select(&swi, `
		SELECT
			search_pages.id,
			search_pages.title,
			search_pages.url,
			search_pages.snippet,
			similarweb_pages.id similarweb_id,
			similarweb_pages.title similarweb_title,
			similarweb_pages.url similarweb_url,
			similarweb_pages.icon_path similarweb_icon,
			similarweb_categories.category similarweb_category
		FROM (
			SELECT
				page_id,
				similarweb_id,
				idf,
				ROW_NUMBER() OVER(PARTITION BY page_id ORDER BY idf DESC) idf_rank
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
		WHERE
			relation.idf_rank <= ?
		UNION
		SELECT
			search_pages.id,
			search_pages.title,
			search_pages.url,
			search_pages.snippet,
			similarweb_pages.id similarweb_id,
			similarweb_pages.title similarweb_title,
			similarweb_pages.url similarweb_url,
			similarweb_pages.icon_path similarweb_icon,
			similarweb_categories.category similarweb_category
		FROM (
			SELECT
				page_id,
				similarweb_id,
				idf,
				ROW_NUMBER() OVER(PARTITION BY page_id ORDER BY idf ASC) idf_rank
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
		WHERE
			relation.idf_rank <= ?
		`, taskId, offset*10, top/2, taskId, offset*10, top-top/2)
	if err != nil {
		return nil, err
	}

	return swi, nil
}

func (s SerpImpl) FetchSerpWithDistributionByID(taskId, offset, top int) ([]models.SerpWithDistributionQueryResult, error) {
	swd := []models.SerpWithDistributionQueryResult{}
	err := s.DB.Select(&swd, `
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
			relation_count.similarweb_count,
			relation_count.category_count / relation_count.similarweb_count category_distribution
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
		return nil, err
	}

	return swd, nil
}
