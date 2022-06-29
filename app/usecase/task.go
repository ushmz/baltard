//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"fmt"
	"ratri/domain/model"
	repo "ratri/domain/repository"
)

// TaskUsecase : Abstract operations that task usecase should have.
type TaskUsecase interface {
	FetchTaskInfo(taskID int) (*model.Task, error)
	CreateTaskAnswer(answer *model.Answer) error
}

// TaskImpl : Struct of task usecase
type TaskImpl struct {
	repository repo.TaskRepository
}

// NewTaskUsecase : Return new task usecase struct
func NewTaskUsecase(taskRepository repo.TaskRepository) TaskUsecase {
	return &TaskImpl{repository: taskRepository}
}

// FetchTaskInfo : Get task information by task ID
func (t *TaskImpl) FetchTaskInfo(taskID int) (*model.Task, error) {
	if t == nil {
		return nil, model.ErrNilReceiver
	}

	task, err := t.repository.FetchTaskInfo(taskID)
	if err != nil {
		return nil, fmt.Errorf("Try to get task information: %w", err)
	}

	return task, nil
}

// CreateTaskAnswer : Create new answer for the task
func (t *TaskImpl) CreateTaskAnswer(answer *model.Answer) error {
	if t == nil {
		return model.ErrNilReceiver
	}

	if err := t.repository.CreateTaskAnswer(answer); err != nil {
		return fmt.Errorf("Try to create new answer: %w", err)
	}

	return nil
}
