package dao

import "github.com/jmoiron/sqlx"

type Condition interface {
	FetchConditionIdByGroupId(groupId int) (int, error)
}

type ConditionImpl struct {
	DB *sqlx.DB
}

func NewCondition(db *sqlx.DB) Condition {
	return &ConditionImpl{DB: db}
}

func (c ConditionImpl) FetchConditionIdByGroupId(groupId int) (int, error) {
	var condition int
	row := c.DB.QueryRow(`
		SELECT
			condition_id
		FROM
			task_condition_relations
		WHERE
			group_id = ?

	`, groupId)

	if err := row.Scan(&condition); err != nil {
		return 0, err
	}

	return condition, nil
}
