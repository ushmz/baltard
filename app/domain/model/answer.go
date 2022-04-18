package model

// Answer : The answer for the task that the participants submit.
type Answer struct {
	// ID : The ID of user.
	ID int `db:"id" json:"id"`

	// UserID : Means external ID.
	UserID int `db:"user_id" json:"user"`

	// TaskID : The identity of task.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : This point out which kind of task did user take.
	ConditionID int `db:"condition_id" json:"condition"`

	// Answer : The Url of evidence of users' decision.
	Answer string `db:"answer" json:"answer"`

	// Reason : The reason of users' decision.
	Reason string `db:"reason" json:"reason"`
}
