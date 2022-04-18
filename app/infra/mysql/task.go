package mysql

import (
	"database/sql"
	"ratri/domain/model"
	repo "ratri/domain/repository"

	"github.com/jmoiron/sqlx"
)

// TaskRepositoryImpl : Implemention of task repository
type TaskRepositoryImpl struct {
	DB *sqlx.DB
}

// NewTaskRepository : Return new task repository struct
func NewTaskRepository(db *sqlx.DB) repo.TaskRepository {
	return &TaskRepositoryImpl{DB: db}
}

// FetchTaskInfo : Fetch task info by task id
func (t *TaskRepositoryImpl) FetchTaskInfo(taskID int) (model.Task, error) {
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
		`, taskID)

	if err := row.StructScan(&task); err != nil {
		if err == sql.ErrNoRows {
			return task, model.NoSuchDataError{}
		}
		return task, err
	}

	return task, nil
}

// UpdateTaskAllocation : Get task ID that the fewest perticipants are allocated
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
	`, gc.Count+1, gc.GroupID)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	// [TODO] How to handle this error?
	tx.Commit()

	return gc.GroupID, nil
}

// GetTaskIDsByGroupID : Get task IDs by group ID
func (t *TaskRepositoryImpl) GetTaskIDsByGroupID(groupID int) ([]int, error) {
	taskIds := []int{}
	err := t.DB.Select(&taskIds, `
		SELECT
			task_id
		FROM
			task_condition_relations
		WHERE
			group_id = ?
	`, groupID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.NoSuchDataError{}
		}
		return nil, err
	}

	return taskIds, nil
}

// GetConditionIDByGroupID : Get condition ID by group ID
func (t *TaskRepositoryImpl) GetConditionIDByGroupID(groupID int) (int, error) {
	var condition int
	row := t.DB.QueryRow(`
		SELECT
			condition_id
		FROM
			task_condition_relations
		WHERE
			group_id = ?

	`, groupID)

	if err := row.Scan(&condition); err != nil {
		if err == sql.ErrNoRows {
			return 0, model.NoSuchDataError{}
		}
		return 0, err
	}

	return condition, nil
}

// CreateTaskAnswer : Create new answer for the task
func (t *TaskRepositoryImpl) CreateTaskAnswer(answer *model.Answer) error {
	_, err := t.DB.NamedExec(`
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
