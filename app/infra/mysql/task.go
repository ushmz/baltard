package mysql

import (
	"database/sql"
	"fmt"
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
	if t == nil {
		return task, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
			return task, fmt.Errorf("Task with given ID(%d) is not found: %w", taskID, model.ErrNoSuchData)
		}
		return task, fmt.Errorf("Try to get task with ID(%d): %w", taskID, err)
	}

	return task, nil
}

// UpdateTaskAllocation : Get task ID that the fewest perticipants are allocated
func (t *TaskRepositoryImpl) UpdateTaskAllocation() (int, error) {
	if t == nil {
		return 0, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
		return 0, fmt.Errorf("Try to get fewest allocated group: %w", err)
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
		return 0, fmt.Errorf("Try to update allocation count with groupID(%d): %w", gc.GroupID, err)
	}

	// [TODO] How to handle this error?
	tx.Commit()

	return gc.GroupID, nil
}

// GetTaskIDsByGroupID : Get task IDs by group ID
func (t *TaskRepositoryImpl) GetTaskIDsByGroupID(groupID int) ([]int, error) {
	if t == nil {
		return []int{}, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
		return nil, fmt.Errorf("Try to get task IDs with group ID(%d): %w", groupID, err)
	}

	return taskIds, nil
}

// GetConditionIDByGroupID : Get condition ID by group ID
func (t *TaskRepositoryImpl) GetConditionIDByGroupID(groupID int) (int, error) {
	if t == nil {
		return 0, fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
			return 0, fmt.Errorf("Condition ID with group ID(%d) is not found: %w", groupID, model.ErrNoSuchData)
		}
		return 0, fmt.Errorf("Try to get condition ID with groupID(%d): %w", groupID, err)
	}

	return condition, nil
}

// CreateTaskAnswer : Create new answer for the task
func (t *TaskRepositoryImpl) CreateTaskAnswer(answer *model.Answer) error {
	if t == nil {
		return fmt.Errorf("Called with nil receiver: %w", model.ErrNilReceiver)
	}

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
		return fmt.Errorf("Try to create new answer with parameters(%v): %w", answer, err)
	}
	return nil
}
