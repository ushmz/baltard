package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"
)

type LinkedPageRepositoryImpl struct {
	DB *sqlx.DB
}

func NewLinkedPageRepository(db *sqlx.DB) repo.LinkedPageRepository {
	return &LinkedPageRepositoryImpl{DB: db}
}

// Get gets a `LinkedPage` records specified with `linkedPageId`.
func (r *LinkedPageRepositoryImpl) Get(linkedPageId int) (model.LinkedPage, error) {
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
	if err := r.DB.Get(&linked, q, linkedPageId); err != nil {
		if err == sql.ErrNoRows {
			return linked, model.NoSuchDataError{}
		}
		return linked, err
	}
	return linked, nil
}

func (r *LinkedPageRepositoryImpl) GetBySearchPageId(pageId, taskId, top int) (*[]model.LinkedPage, error) {
	linked := []model.LinkedPage{}

	q := `
		SELECT
			sp.id,
			sp.title,
			sp.url,
			sp.icon_path,
			sc.category 
		FROM 
			similarweb_pages sp
		JOIN
			similarweb_categories sc
		ON
			sp.category = sc.id
		WHERE
			sp.id IN (
				SELECT * FROM (
					SELECT
						similarweb_id
					FROM
						search_page_similarweb_relation
					WHERE
						page_id = ?
					AND
						task_id = ?
					ORDER BY
						idf DESC
					LIMIT
						? ) linked
			)
	`

	if err := r.DB.Select(&linked, q, pageId, taskId, top); err != nil {
		return &linked, err
	}

	return &linked, nil
}

// Select gets listed `LinkedPage` specified with argument `linkedPageIds`.
// [TODO] Which is better?
// - Take only `[]int` argument and cast it to `[]interface{}`.
// - It implicitly assume that passed argument `linkedPageIds` is list of number
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
