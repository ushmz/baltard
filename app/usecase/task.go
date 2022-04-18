//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"ratri/domain/model"
	repo "ratri/domain/repository"
)

// TaskUsecase : Abstract operations that task usecase should have.
type TaskUsecase interface {
	FetchTaskInfo(taskID int) (model.Task, error)
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
func (t *TaskImpl) FetchTaskInfo(taskID int) (model.Task, error) {
	return t.repository.FetchTaskInfo(taskID)
}

// CreateTaskAnswer : Create new answer for the task
func (t *TaskImpl) CreateTaskAnswer(answer *model.Answer) error {
	return t.repository.CreateTaskAnswer(answer)
}
