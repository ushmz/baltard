package model

// UserParam : Struct for request of `/signup` endpoint
type UserParam struct {
	// Uid : User name/ID for label.
	Uid string `json:"uid"`
}

// User : Struct for user information.
type User struct {
	// Id : The ID of user.
	Id int `db:"id" json:"id"`

	// Uid : External user_id (like crowdsourcing site).
	Uid string `db:"uid" json:"uid"`

	// Token :
	Token string `json:"token"`
}

// UserResponse : Struct for response body of `CreateUser` handler
type UserResponse struct {
	// Token : Bearer token for authentication
	Token string `json:"token"`

	// UserId : Unique ID used in DB.
	UserId int `json:"user"`

	// TaskIds : Shows the IDs that user perform
	TaskIds []int `json:"tasks"`

	// ConditionId : Assigned condition ID
	ConditionId int `json:"condition"`

	// GroupId : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupId int `json:"group"`
}
