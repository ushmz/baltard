package models

// TaskTimeLogParam : Struct for task viewing time log request body
type TaskTimeLogParam struct {
	// Id : The ID of each log record.
	Id string `db:"id" json:"id"`

	// Uid : The ID of user (worker)
	Uid string `db:"uid" json:"uid"`

	// TimeOnPage : User's page viewing time.
	TimeOnPage int64 `db:"time_on_page" json:"time"`

	// Url : The url of user viewing page.
	Url string `db:"url" json:"url"`

	// TaskId : The Id of task that user working.
	TaskId int64 `db:"task_id" json:"task"`

	// ConditionId : User's condition Id that means group and task category.
	ConditionId int64 `db:"condition_id" json:"condition"`
}

// SearchPageClickLogParam : Create page click log. Table name is `behavior_logs_click`.
type SearchPageClickLogParamWithoutVisible struct {
	// Id : The ID of each log record.
	Id string `db:"id" json:"id"`

	// Uid : The ID of user (worker)
	Uid string `db:"uid" json:"uid"`

	// TaskId : The Id of task that user working.
	TaskId int64 `db:"task_id" json:"taskId"`

	// ConditionId : User's condition Id that means group and task category.
	ConditionId int64 `db:"condition_id" json:"conditionId"`

	// Time : User's page viewing time.
	Time int64 `db:"time_on_page" json:"time"`

	// Page : The Id of page that user clicked.
	Page int64 `db:"serp_page" json:"page"`

	// Rank : Search result rank that user clicked.
	Rank int64 `db:"serp_rank" json:"rank"`
}

// SearchPageClickLogParamWithVisible : Create page click log. Table name is `behavior_logs_click`.
type SearchPageClickLogParam struct {
	// Id : The ID of each log record.
	Id string `db:"id" json:"id"`

	// Uid : The ID of user (worker)
	User int `db:"user_id" json:"user"`

	// TaskId : The Id of task that user working.
	TaskId int `db:"task_id" json:"taskId"`

	// ConditionId : User's condition Id that means group and task category.
	ConditionId int `db:"condition_id" json:"conditionId"`

	// Time : User's page viewing time.
	Time int `db:"time_on_page" json:"time"`

	// Page : The Id of page that user clicked.
	Page int `db:"serp_page" json:"page"`

	// Rank : Search result rank that user clicked.
	Rank int `db:"serp_rank" json:"rank"`

	// IsVisible : Risk is visible or not.
	IsVisible bool `db:"is_visible" json:"visible"`
}
