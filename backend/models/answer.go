package models

type Answer struct {
	Id                       string `db:"id" json:"id"`
	Uid                      string `db:"uid" json:"uid"`
	TaskId                   int64  `db:"task_id" json:"taskId"`
	ConditionId              int64  `db:"condition_id" json:"conditionId"`
	TaskConditionRelationsId int64  `db:"task_condition_relations_id" json:"taskConditionRelationsId"`
	AuthorId                 int64  `db:"author_id" json:"authorId"`
	HotelId                  int64  `db:"hotel_id" json:"hotelId"`
	Reason                   string `db:"reason" json:"reason"`
}
