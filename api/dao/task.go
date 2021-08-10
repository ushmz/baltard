package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/ymmt3-lab/koolhaas/backend/api/models"
)

type Task interface {
	AllocateTask() (int, error)
	FetchTaskIdsByGroupId(groupId int) ([]int, error)
	FetchTaskInfo(taskId int) (*models.Task, error)
}

type TaskImpl struct {
	DB *sqlx.DB
}

func NewTask(db *sqlx.DB) Task {
	return &TaskImpl{DB: db}
}

func (t TaskImpl) AllocateTask() (int, error) {
	tx := t.DB.MustBegin()
	gc := models.GroupCounts{}
	err := tx.Get(&gc, `
		SELECT
			group_id,
			`+"`count`"+`
		FROM
			group_counts
		WHERE`+
		"   `count`"+` = (
				SELECT
					MIN(`+"`count`"+`)
				FROM
					group_counts
			)
		LIMIT
			1
	`)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(`
		UPDATE
			group_counts
		SET`+
		"   `count`"+` = ?
		WHERE
			group_id = ?
	`, gc.Count+1, gc.GroupId)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	// [TODO] How to handle this error?
	tx.Commit()

	return gc.GroupId, nil
}

func (t TaskImpl) FetchTaskIdsByGroupId(groupId int) ([]int, error) {
	taskIds := []int{}
	err := t.DB.Select(&taskIds, `
		SELECT
			task_id
		FROM
			task_condition_relations
		WHERE
			group_id = ?
	`, groupId)

	if err != nil {
		return nil, err
	}

	return taskIds, nil
}

// [TODO] Difference of return value
// FetchTaskInfo : Fetch task info by task id
func (t TaskImpl) FetchTaskInfo(taskId int) (*models.Task, error) {
	task := models.Task{}
	row := t.DB.QueryRowx(`
		SELECT
			id,
			query,
			title,
			description,
			search_url
		FROM
			tasks
		WHERE
			id = ?
		`, taskId)

	if err := row.StructScan(&task); err != nil {
		return nil, err
	}

	return &task, nil
}
