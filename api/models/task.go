package models

type Task struct {
	// Id : The ID of task
	Id int `db:"id" json:"id"`

	// Query : Search query for this task.
	Query string `db:"query" json:"query"`

	// Title : Title of this task.
	Title string `db:"title" json:"title"`

	// Description : Description text of task.
	Description string `db:"description" json:"description"`

	// AuthorId : Author of this task.
	// AuthorId string `db:"author_id" json:"authorId"`

	// SearchUrl : Url used in this task.
	SearchUrl string `db:"search_url" json:"searchUrl"`

	// Type : Task type, used like category if needed.
	// Type sql.NullString `db:"type" json:"type"`
}

type TaskAnswer struct {
	// Id : The ID of user.
	Id int `db:"id" json:"id"`

	// UserId : Means external Id.
	UserId int `db:"user_id" json:"user"`

	// TaskId : The identity of task.
	TaskId int `db:"task_id" json:"task"`

	// ConditionId : This point out which kind of task did user take.
	ConditionId int `db:"condition_id" json:"condition"`

	// Answer : The Url of evidence of users' decision.
	Answer string `db:"answer" json:"answer"`

	// Reason : The reason of users' decision.
	Reason string `db:"reason" json:"reason"`
}

// GroupCounts : Struct for group count
type GroupCounts struct {
	GroupId int `db:"group_id" json:"groupId"`
	Count   int `db:"count" json:"count"`
}
