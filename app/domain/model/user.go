package model

// UserParam : Struct for request of `/signup` endpoint
type UserParam struct {
	// UID : User name/ID for label.
	UID string `json:"uid"`
}

// User : Struct for user information.
type User struct {
	// ID : The ID of user.
	ID int `db:"id" json:"id"`

	// UID : External user_id (like crowdsourcing site).
	UID string `db:"uid" json:"uid"`

	// Token :
	Token string `json:"token"`
}

// UserResponse : Struct for response body of `CreateUser` handler
type UserResponse struct {
	// Token : Bearer token for authentication
	Token string `json:"token"`

	// UserID : Unique ID used in DB.
	UserID int `json:"user"`

	// TaskIDs : Shows the IDs that user perform
	TaskIDs []int `json:"tasks"`

	// ConditionID : Assigned condition ID
	ConditionID int `json:"condition"`

	// GroupID : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupID int `json:"group"`
}
