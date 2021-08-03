package models

import "database/sql"

type Task struct {
	// Id : The ID of task
	Id int64 `db:"id" json:"id"`

	// Query : Search query for this task.
	Query string `db:"query" json:"query"`

	// Title : Title of this task.
	Title string `db:"title" json:"title"`

	// Description : Description text of task.
	Description string `db:"description" json:"description"`

	// AuthorId : Author of this task.
	AuthorId string `db:"author_id" json:"authorId"`

	// SearchUrl : Url used in this task.
	SearchUrl string `db:"search_url" json:"searchUrl"`

	// Type : Task type, used like category if needed.
	Type sql.NullString `db:"type" json:"type"`
}

type TaskAnswer struct {
	// Id : The ID of user.
	Id int `db:"id" json:"id"`

	// Uid : Means external Id.
	Uid string `db:"uid" json:"uid"`

	// TaskId : The identity of task.
	TaskId int `db:"task_id" json:"task"`

	// ConditionId : This point out which kind of task did user take.
	ConditionId int `db:"condition_id" json:"condition"`

	// AuthorId : The identity of task author
	AuthorId int `db:"author_id" json:"authorId"`

	// Answer : The Url of evidence of users' decision.
	Answer string `db:"answer" json:"answer"`

	// Reason : The reason of users' decision.
	Reason string `db:"reason" json:"reason"`
}
