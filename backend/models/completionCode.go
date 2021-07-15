package models

type CompletionCode struct {
	// CompletionCode : Completion code to confirm task completion for workers.
	CompletionCode int64 `db:"completion_code" json:"completionCode"`
}
