package model

import "time"

// SerpViewingLogParam : Struct for task viewing time log request body
type SerpViewingLogParam struct {
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionId : User's condition ID that means group and task category.
	ConditionId int `db:"condition_id" json:"condition"`
}

type SerpViewingLog struct {
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionId : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// DwellTime : Time(sec.) that the user spend in SERP
	DwellTime int `db:"time_on_page" json:"dwell_time"`

	// CreatedAt :
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	// UpdatedAt :
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// PageViewingLogParam : Struct for each search result page viewing time log request body
type PageViewingLogParam struct {
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// PageID : Page ID that user view.
	PageID int `db:"page_id" json:"page"`
}

type PageViewingLog struct {
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

	// CreatedAt :
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	// UpdatedAt :
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
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

	// CreatedAt :
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	// UpdatedAt :
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
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

type SearchSession struct {
	// UserID : Assigned ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// StartedAt :
	StartedAt time.Time `db:"started_at" json:"started_at"`

	// EndedAt :
	EndedAt time.Time `db:"ended_at" json:"ended_at"`
}
