package usecase

import (
	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"
)

type Log interface {
	StoreTaskTimeLog(*model.TaskTimeLogParamWithTime) error
	CumulateTaskTimeLog(*model.TaskTimeLogParam) error
	StoreSerpClickLog(*model.SearchPageClickLogParam) error
	StoreSearchSeeion(*model.SearchSession) error
}

type LogImpl struct {
	repository repo.LogRepository
}

func NewLogUsecase(logRepository repo.LogRepository) Log {
	return &LogImpl{repository: logRepository}
}

func (l *LogImpl) StoreTaskTimeLog(p *model.TaskTimeLogParamWithTime) error {
	return l.repository.StoreTaskTimeLog(p)
}

func (l *LogImpl) CumulateTaskTimeLog(p *model.TaskTimeLogParam) error {
	return l.repository.CumulateTaskTimeLog(p)
}

func (l *LogImpl) StoreSerpClickLog(p *model.SearchPageClickLogParam) error {
	return l.repository.StoreSerpClickLog(p)
}

func (l *LogImpl) StoreSearchSeeion(p *model.SearchSession) error {
	return l.repository.StoreSearchSeeion(p)
}
