package models

// UserParam : Struct for request of `/signup` endpoint
type UserParam struct {
	// Uid : User name/Id for label.
	Uid string `json:"uid"`
}

// User : Struct for `users` table model
type User struct {
	// Id : The ID of user.
	Id int64 `db:"id" json:"id"`

	// Uid : User name/Id for label.
	Uid string `db:"uid" json:"uid"`

	// ExternalId : External user Id.
	ExternalId string `db:"external_id" json:"externalId"`

	// Email : uid + dummy domain use for firebase authentication.
	Email string `db:"email" json:"email"`

	// CreatedAt : Auto generated datatime information.
	CreatedAt string `db:"created_at" json:"createdAt"`
}

type ExistUser struct {
	// Id : The ID of user.
	Id int64 `db:"id" json:"id"`

	// Uid : User name/Id for label.
	Uid string `db:"uid" json:"uid"`

	// Secret : Generated secret string.
	Secret string `db:"generated_secret" json:"secret"`
}

// UserResponse : Struct for response body of `CreateUser` handler
type UserResponse struct {
	// Exist : Given uid is exist on DB ot not.
	Exist bool `json:"exist"`

	// UserId : Unique Id used in koolhaas DB.
	UserId int64 `json:"user"`

	// Secret : Generated secret string.
	Secret string `json:"secret"`

	// TaskIds : Allocated tasks Ids.
	TaskIds []int `json:"tasks"`

	// ConditionId : Allocated condition Id.
	ConditionId int `json:"condition"`

	// GroupId : Task and condition distinction.
	GroupId int `json:"group"`
}
