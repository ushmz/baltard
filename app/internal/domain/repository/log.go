package repository

import (
	"ratri/internal/domain/model"
)

type LogRepository interface {
	StoreTaskTimeLog(*model.TaskTimeLogParamWithTime) error
	CumulateTaskTimeLog(*model.TaskTimeLogParam) error
	StoreSerpClickLog(*model.SearchPageClickLogParam) error
	StoreSearchSeeion(*model.SearchSession) error
}
