package model

import "time"

// SerpDwellTimeLogParam : Struct for task viewing time log request body
type SerpDwellTimeLogParam struct {
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`
}

// SerpDwellTimeLog : This shows how long each participant views the SERP.
type SerpDwellTimeLog struct {
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionId : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// DwellTime : Time(sec.) that the user spend in SERP.
	DwellTime int `db:"time_on_page" json:"dwell_time"`

	// Created : Timestamp that this record is created.
	Created time.Time `db:"created_at" json:"created_at"`

	// Updated : Timestamp that this record is updated.
	Updated time.Time `db:"updated_at" json:"updated_at"`
}

// PageDwellTimeLogParam : Struct for each search result page viewing time log request body
type PageDwellTimeLogParam struct {
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// PageID : Page ID that user view.
	PageID int `db:"page_id" json:"page"`
}

// PageDwellTimeLog : This shows how long each participant views the result page.
type PageDwellTimeLog struct {
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// PageID : Page ID that user view.
	PageID int `db:"page_id" json:"page"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// DwellTime : Time(sec.) that the user spend in SERP
	DwellTime int `db:"time_on_page" json:"dwell_time"`

	// Created : Timestamp that this record is created.
	Created time.Time `db:"created_at" json:"created_at"`

	// Updated : Timestamp that this record is updated.
	Updated time.Time `db:"updated_at" json:"updated_at"`
}

// SearchPageEventLogParam : Struct for page click log request body.
type SearchPageEventLogParam struct {
	// ID : The ID of each log record.
	ID string `db:"id" json:"id"`

	// Uid : The ID of user (worker)
	User int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// Time : User's page viewing time.
	Time int `db:"time_on_page" json:"time"`

	// Page : The ID of page that user clicked.
	Page int `db:"serp_page" json:"page"`

	// Rank : Search result rank that user clicked.
	Rank int `db:"serp_rank" json:"rank"`

	// IsVisible : Risk is visible or not.
	IsVisible bool `db:"is_visible" json:"visible"`

	// Event : It is expected to be "click", "hover" or "paginate"
	Event string `db:"event" json:"event"`
}

// SearchPageEventLog : This shows when each participant take an action.
type SearchPageEventLog struct {
	// ID : The ID of each log record.
	ID string `db:"id" json:"id"`

	// UserId : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskId : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionId : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// Time : User's page viewing time.
	Time int `db:"time_on_page" json:"time"`

	// Page : The ID of page that user clicked.
	Page int `db:"serp_page" json:"page"`

	// Rank : Search result rank that user clicked.
	Rank int `db:"serp_rank" json:"rank"`

	// IsVisible : Risk is visible or not.
	IsVisible bool `db:"is_visible" json:"visible"`

	// Event : It is expected to be "click", "hover" or "paginate"
	Event string `db:"event" json:"event"`

	// Created : Timestamp that this record is created.
	Created time.Time `db:"created_at" json:"created_at"`

	// Updated : Timestamp that this record is updated.
	Updated time.Time `db:"updated_at" json:"updated_at"`
}

// SearchSessionParam : Struct fot search session request body.
type SearchSessionParam struct {
	// UserID : Assigned ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`
}

// SearchSession : How long each participant takes for the task.
type SearchSession struct {
	// UserID : Assigned ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// Started : Timestamp that the participant starts the task.
	Started time.Time `db:"started_at" json:"started_at"`

	// Ended : Timestamp that the participant ends the task.
	Ended time.Time `db:"ended_at" json:"ended_at"`
}
