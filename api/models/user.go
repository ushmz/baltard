package models

// UserParam : Struct for request of `/signup` endpoint
type UserParam struct {
	// Uid : User name/Id for label.
	Uid string `json:"uid"`
}

// User :
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

	// UserId : Unique Id used in DB.
	UserId int `json:"user"`

	// Secret : Generated secret string.
	Secret string `json:"secret"`

	// TaskIds : Allocated tasks Ids.
	TaskIds []int `json:"tasks"`

	// ConditionId : Allocated condition Id.
	ConditionId int `json:"condition"`

	// GroupId : Task and condition distinction.
	GroupId int `json:"group"`
}
