//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"ratri/domain/model"
)

// TaskRepository : Abstract operations that `Task` model should have.
type TaskRepository interface {
	FetchTaskInfo(taskID int) (*model.Task, error)
	CreateTaskAnswer(*model.Answer) error
	UpdateTaskAllocation() (int, error)
	// Think about creating `ConditionReporitory` and `GroupReporitory` ???
	GetTaskIDsByGroupID(groupID int) ([]int, error)
	GetConditionIDByGroupID(groupID int) (int, error)
}
