package dao

import (
	"baltard/api/model"

	"github.com/jmoiron/sqlx"
)

type Task interface {
	FetchTaskInfo(taskId int) (*model.Task, error)
	AllocateTask() (int, error)
	FetchTaskIdsByGroupId(groupId int) ([]int, error)
	FetchConditionIdByGroupId(groupId int) (int, error)
	SubmitTaskAnswer(*model.Answer) error
}

type TaskImpl struct {
	DB *sqlx.DB
}

func NewTask(db *sqlx.DB) Task {
	return &TaskImpl{DB: db}
}

// FetchTaskInfo : Fetch task info by task id
func (t TaskImpl) FetchTaskInfo(taskId int) (*model.Task, error) {
	task := model.Task{}
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

func (t TaskImpl) AllocateTask() (int, error) {
	tx := t.DB.MustBegin()
	gc := model.GroupCounts{}
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

func (t TaskImpl) FetchConditionIdByGroupId(groupId int) (int, error) {
	var condition int
	row := t.DB.QueryRow(`
		SELECT
			condition_id
		FROM
			task_condition_relations
		WHERE
			group_id = ?

	`, groupId)

	if err := row.Scan(&condition); err != nil {
		return 0, err
	}

	return condition, nil
}

func (a TaskImpl) SubmitTaskAnswer(answer *model.Answer) error {
	_, err := a.DB.NamedExec(`
		INSERT INTO
			answers (
				user_id,
				task_id,
				condition_id,
				answer,
				reason
			)
		VALUES (
			:user_id,
			:task_id,
			:condition_id,
			:answer,
			:reason
		)`, answer)
	if err != nil {
		return err
	}
	return nil
}
