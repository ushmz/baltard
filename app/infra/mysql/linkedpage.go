package mysql

import (
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"

	"ratri/domain/model"
	repo "ratri/domain/repository"
)

// LinkedPageRepositoryImpl : Implemention of LinkedPage repository
type LinkedPageRepositoryImpl struct {
	DB *sqlx.DB
}

// NewLinkedPageRepository : Return new LinkedPage repository struct
func NewLinkedPageRepository(db *sqlx.DB) repo.LinkedPageRepository {
	return &LinkedPageRepositoryImpl{DB: db}
}

// Get gets a `LinkedPage` records specified with `linkedPageId`.
func (r *LinkedPageRepositoryImpl) Get(linkedPageID int) (model.LinkedPage, error) {
	linked := model.LinkedPage{}

	q := `
		SELECT
			s.id id,
			s.title title,
			s.url url,
			s.icon_path icon_path,
			sc.category
		FROM
			similarweb_pages s
		LEFT JOIN
			similarweb_categories sc ON s.category = sc.id
		WHERE
			s.id = ?
	`
	if err := r.DB.Get(&linked, q, linkedPageID); err != nil {
		if err == sql.ErrNoRows {
			return linked, model.NoSuchDataError{}
		}
		return linked, err
	}
	return linked, nil
}

// GetBySearchPageIDs : Get linked pages for the Icon UI by given search page IDs
func (r *LinkedPageRepositoryImpl) GetBySearchPageIDs(pageID []int, taskID, top int) (*[]model.SearchPageWithLinkedPage, error) {
	linked := []model.SearchPageWithLinkedPage{}

	q := `
		SELECT
			rel.page_id,
			sp.id,
			sp.title,
			sp.url,
			sp.icon_path,
			sc.category
		FROM (
			SELECT
				*
			FROM (
				SELECT
					page_id,
					similarweb_id,
					ROW_NUMBER() OVER (PARTITION BY page_id ORDER BY idf DESC) idf_rank
				FROM
					search_page_similarweb_relation
				WHERE
					page_id IN( ?` + strings.Repeat(", ?", len(pageID)-1) + `)
					AND task_id = ?
				ORDER BY
					page_id ASC
			) linked
			WHERE
				idf_rank <= ?) rel
		LEFT JOIN similarweb_pages sp ON rel.similarweb_id = sp.id
		LEFT JOIN similarweb_categories sc ON sp.category = sc.id;
	`

	a := []interface{}{}
	for _, v := range pageID {
		a = append(a, v)
	}
	a = append(a, taskID)
	a = append(a, top)

	if err := r.DB.Select(&linked, q, a...); err != nil {
		return &linked, err
	}

	return &linked, nil
}

// GetRatioBySearchPageIDs : Get Ratio information for the Ratio UI by given search page IDs
func (r *LinkedPageRepositoryImpl) GetRatioBySearchPageIDs(pageIds []int, taskID int) (*[]model.SearchPageWithLinkedPageRatio, error) {
	linked := []model.SearchPageWithLinkedPageRatio{}

	q := `
		SELECT DISTINCT
			rel.page_id,
			sc.category,
			COUNT(*) OVER(PARTITION BY rel.page_id, sp.category) category_count
		FROM
			search_page_similarweb_relation rel
		LEFT JOIN similarweb_pages sp ON rel.similarweb_id = sp.id
		LEFT JOIN similarweb_categories sc ON sp.category = sc.id
		WHERE
			task_id = ?
		AND
			page_id IN ( ? ` + strings.Repeat(", ?", len(pageIds)-1) + ` )
		ORDER BY
			page_id, category_count DESC
	`

	a := []interface{}{}
	a = append(a, taskID)
	for _, v := range pageIds {
		a = append(a, v)
	}

	if err := r.DB.Select(&linked, q, a...); err != nil {
		return &linked, err
	}

	return &linked, nil
}

// Select gets listed `LinkedPage` specified with argument `linkedPageIds`.
// [TODO] Which is better?
// - Take only `[]int` argument and cast it to `[]interface{}`.
// - It implicitly assume that passed argument `linkedPageIDs` is list of number
//   (or check argument type and if it's not int value, return error)
//   and make argument type as `[]interface{}`
func (r *LinkedPageRepositoryImpl) Select(linkedPageIds []int) (*[]model.LinkedPage, error) {
	linked := []model.LinkedPage{}

	dest := []interface{}{}
	for _, v := range linkedPageIds {
		dest = append(dest, v)
	}

	q := `
		SELECT
			s.id id,
			s.title title,
			s.url url,
			s.icon_path icon_path,
			sc.category
		FROM
			similarweb_pages s
		LEFT JOIN
			similarweb_categories sc ON s.category = sc.id
		WHERE
			s.id IN (?)
	`
	q, a, err := sqlx.In(q, dest)
	if err != nil {
		return &linked, err
	}

	if err := r.DB.Select(&linked, q, a...); err != nil {
		if err == sql.ErrNoRows {
			return &linked, model.NoSuchDataError{}
		}
		return &linked, err
	}
	return &linked, nil
}

// List gets all `LinkedPage` records from DB
func (r *LinkedPageRepositoryImpl) List(offset, limit int) (*[]model.LinkedPage, error) {
	linked := []model.LinkedPage{}

	q := `
		SELECT
			s.id id,
			s.title title,
			s.url url,
			s.icon_path icon_path,
			sc.category
		FROM
			similarweb_pages s
		LEFT JOIN
			similarweb_categories sc ON s.category = sc.id
		LIMIT
			?, ?
	`
	if err := r.DB.Select(&linked, q, offset, limit); err != nil {
		if err == sql.ErrNoRows {
			return &linked, model.NoSuchDataError{}
		}
		if len(linked) == 0 {
			return &linked, model.NoSuchDataError{}
		}
		return &linked, err
	}
	return &linked, nil
}

// Create creates new `LinkedPage` record.
// However it should not be created by API, so this implementation is empty.
// func (r LinkedPageRepositoryImpl) Create(model.LinkedPage) (model.LinkedPage, error)

// Update updates a `LinkedPage` record.
// However it should not be updated by API, so this implementation is empty.
// func (r LinkedPageRepositoryImpl) Update(model.LinkedPage) (model.LinkedPage, error)

// Delete deletes a `LinkedPage` record specified with `linkedPageId`.
// However it should not be created by API, so this implementation is empty.
// func (r LinkedPageRepositoryImpl) Delete(linkedPageId int) error
