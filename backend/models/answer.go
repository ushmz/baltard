package models

type Answer struct {
	Id          string `db:"id" json:"id"`
	Uid         string `db:"uid" json:"uid"`
	TaskId      int64  `db:"task_id" json:"task"`
	ConditionId int64  `db:"condition_id" json:"condition"`
	AuthorId    int64  `db:"author_id" json:"author"`
	Reason      string `db:"reason" json:"reason"`
}
