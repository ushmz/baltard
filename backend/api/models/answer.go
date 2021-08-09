package models

type Answer struct {
	Id          string `db:"id" json:"id"`
	Uid         string `db:"uid" json:"uid"`
	TaskId      int    `db:"task_id" json:"task"`
	ConditionId int    `db:"condition_id" json:"condition"`
	AuthorId    int    `db:"author_id" json:"author"`
	Reason      string `db:"reason" json:"reason"`
}
