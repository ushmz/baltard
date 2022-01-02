//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"ratri/internal/domain/model"
)

type TaskRepository interface {
	FetchTaskInfo(taskId int) (model.Task, error)
	CreateTaskAnswer(*model.Answer) error
	UpdateTaskAllocation() (int, error)
	// Think about creating `ConditionReporitory` and `GroupReporitory` ???
	GetTaskIdsByGroupId(groupId int) ([]int, error)
	GetConditionIdByGroupId(groupId int) (int, error)
}
