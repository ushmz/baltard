package repository

import (
	"baltard/internal/domain/model"
)

type SerpRepository interface {
	FetchSerpByTaskID(taskId, offset int) ([]model.SearchPage, error)
	FetchSerpWithIconByTaskID(taskId, offset, top int) ([]model.SerpWithIconQueryResult, error)
	FetchSerpWithRatioByTaskID(taskId, offset, top int) ([]model.SerpWithRatioQueryResult, error)
}
