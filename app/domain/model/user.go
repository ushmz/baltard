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

	// Secret : Generated secret string.
	Secret string `db:"generated_secret" json:"secret"`
}

// UserResponse : Struct for response body of `CreateUser` handler
type UserResponse struct {
	// Exist : Given uid is exist on DB ot not.
	Exist bool `json:"exist"`

	// UserId : Unique ID used in DB.
	UserId int `json:"user"`

	// Secret : Generated secret string.
	Secret string `json:"secret"`

	// TaskIds : Shows the IDs that user perform
	TaskIds []int `json:"tasks"`

	// ConditionId : Assigned condition ID
	ConditionId int `json:"condition"`

	// GroupId : The ID assigned to the pair of "task IDs" and "condition ID"
	GroupId int `json:"group"`
}
