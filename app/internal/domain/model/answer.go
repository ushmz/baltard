package model

type Answer struct {
	// Id : The ID of user.
	Id int `db:"id" json:"id"`

	// UserId : Means external ID.
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
