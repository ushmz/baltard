package service

import (
	"baltard/api/dao"
	"baltard/api/model"
)

type Task interface {
	FetchTaskInfo(taskId int) (*model.Task, error)
	SubmitTaskAnswer(answer *model.Answer) error
}

type TaskImpl struct {
	taskDao dao.Task
}

func NewTaskService(taskDao dao.Task) Task {
	return &TaskImpl{taskDao: taskDao}
}

func (t *TaskImpl) FetchTaskInfo(taskId int) (*model.Task, error) {
	return t.taskDao.FetchTaskInfo(taskId)
}

func (t *TaskImpl) SubmitTaskAnswer(answer *model.Answer) error {
	return t.taskDao.SubmitTaskAnswer(answer)
}
