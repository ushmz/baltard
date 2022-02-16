package mysql

import (
	"database/sql"
	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"

	"github.com/jmoiron/sqlx"
)

type TaskRepositoryImpl struct {
	DB *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) repo.TaskRepository {
	return &TaskRepositoryImpl{DB: db}
}

// FetchTaskInfo : Fetch task info by task id
func (t *TaskRepositoryImpl) FetchTaskInfo(taskId int) (model.Task, error) {
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
		if err == sql.ErrNoRows {
			return task, model.NoSuchDataError{}
		}
		return task, err
	}

	return task, nil
}

func (t *TaskRepositoryImpl) UpdateTaskAllocation() (int, error) {
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

func (t *TaskRepositoryImpl) GetTaskIdsByGroupId(groupId int) ([]int, error) {
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
		if err == sql.ErrNoRows {
			return nil, model.NoSuchDataError{}
		}
		return nil, err
	}

	return taskIds, nil
}

func (t *TaskRepositoryImpl) GetConditionIdByGroupId(groupId int) (int, error) {
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
		if err == sql.ErrNoRows {
			return 0, model.NoSuchDataError{}
		}
		return 0, err
	}

	return condition, nil
}

func (a *TaskRepositoryImpl) CreateTaskAnswer(answer *model.Answer) error {
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
