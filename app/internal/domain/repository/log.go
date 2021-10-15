package repository

import (
	"ratri/internal/domain/model"
)

type LogRepository interface {
	StoreTaskTimeLog(*model.TaskTimeLogParam) error
	CumulateTaskTimeLog(*model.TaskTimeLogParam) error
	StoreSerpClickLog(*model.SearchPageClickLogParam) error
	StoreSearchSeeion(*model.SearchSession) error
}
