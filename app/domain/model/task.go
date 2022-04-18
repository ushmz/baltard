package model

// Task : Struct for Task information.
type Task struct {
	// ID : The ID of task
	ID int `db:"id" json:"id"`

	// Query : Search query for this task.
	Query string `db:"query" json:"query"`

	// Title : Title of this task.
	Title string `db:"title" json:"title"`

	// Description : Description text of task.
	Description string `db:"description" json:"description"`

	// SearchURL : Url used in this task.
	SearchURL string `db:"search_url" json:"searchUrl"`
}

// GroupCounts : Struct for group count
type GroupCounts struct {
	// GroupID : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupID int `db:"group_id" json:"groupId"`

	// Count : Shows how many users are assigned to this group.
	Count int `db:"count" json:"count"`
}

// TaskInfo : Struct for response of which task is assigned.
type TaskInfo struct {
	// GroupID : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupID int

	// ConditionID : Assigned condition ID
	ConditionID int

	// TaskIDs : Shows the IDs that user perform
	TaskIDs []int
}
