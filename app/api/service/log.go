package service

import (
	"baltard/api/dao"
	"baltard/api/model"
)

type Log interface {
	CreateTaskTimeLog(*model.TaskTimeLogParam) error
	CreateSerpClickLog(*model.SearchPageClickLogParam) error
	StoreSearchSeeion(*model.SearchSession) error
}

type LogImpl struct {
	dao dao.Log
}

func NewLogService(logDao dao.Log) Log {
	return &LogImpl{dao: logDao}
}

func (l *LogImpl) CreateTaskTimeLog(p *model.TaskTimeLogParam) error {
	return l.dao.CreateTaskTimeLog(p)
}

func (l *LogImpl) CreateSerpClickLog(p *model.SearchPageClickLogParam) error {
	return l.dao.CreateSerpClickLog(p)
}

func (l *LogImpl) StoreSearchSeeion(p *model.SearchSession) error {
	return l.dao.StoreSearchSeeion(p)
}
