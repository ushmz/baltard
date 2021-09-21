package usecase

import (
	"baltard/internal/domain/model"
	repo "baltard/internal/domain/repository"
)

type Task interface {
	FetchTaskInfo(taskId int) (*model.Task, error)
	CreateTaskAnswer(answer *model.Answer) error
}

type TaskImpl struct {
	repository repo.TaskRepository
}

func NewTaskUsecase(taskRepository repo.TaskRepository) Task {
	return &TaskImpl{repository: taskRepository}
}

func (t *TaskImpl) FetchTaskInfo(taskId int) (*model.Task, error) {
	return t.repository.FetchTaskInfo(taskId)
}

func (t *TaskImpl) CreateTaskAnswer(answer *model.Answer) error {
	return t.repository.CreateTaskAnswer(answer)
}
