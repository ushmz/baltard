package models

type CompletionCode struct {
	// CompletionCode : Completion code to confirm task completion for workers.
	CompletionCode int `db:"completion_code" json:"completionCode"`
}
