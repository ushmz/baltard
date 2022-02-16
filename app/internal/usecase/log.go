//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"bytes"
	"errors"
	"ratri/internal/domain/model"
	repo "ratri/internal/domain/repository"
	"ratri/internal/domain/store"
)

type LogUsecase interface {
	CumulateSerpViewingTime(model.SerpViewingLogParam) error
	CumulatePageViewingTime(model.PageViewingLogParam) error
	StoreSerpEventLog(model.SearchPageEventLogParam) error
	StoreSearchSeeion(model.SearchSessionParam) error

	ExportSerpViewingTimeLog(bool, store.FileType) (*bytes.Buffer, error)
	ExportPageViewingTimeLog(bool, store.FileType) (*bytes.Buffer, error)
	ExportSerpEventLog(bool, store.FileType) (*bytes.Buffer, error)
	ExportSearchSeeion(bool, store.FileType) (*bytes.Buffer, error)
}

type LogImpl struct {
	repository repo.LogRepository
	store      store.LogStore
}

func NewLogUsecase(logRepository repo.LogRepository, store store.LogStore) LogUsecase {
	return &LogImpl{repository: logRepository, store: store}
}

func (l *LogImpl) CumulateSerpViewingTime(p model.SerpViewingLogParam) error {
	if l == nil {
		return errors.New("LogImpl is nil")
	}
	return l.repository.CumulateSerpViewingTime(p)
}

func (l *LogImpl) CumulatePageViewingTime(p model.PageViewingLogParam) error {
	if l == nil {
		return errors.New("LogImpl is nil")
	}
	return l.repository.CumulatePageViewingTime(p)
}

func (l *LogImpl) StoreSerpEventLog(p model.SearchPageEventLogParam) error {
	if l == nil {
		return errors.New("LogImpl is nil")
	}
	return l.repository.StoreSerpEventLog(p)
}

func (l *LogImpl) StoreSearchSeeion(p model.SearchSessionParam) error {
	if l == nil {
		return errors.New("LogImpl is nil")
	}
	return l.repository.StoreSearchSeeion(p)
}

func (l *LogImpl) ExportSerpViewingTimeLog(header bool, filetype store.FileType) (*bytes.Buffer, error) {
	if l == nil {
		return nil, errors.New("LogImpl is nil")
	}

	data, err := l.repository.FetchAllSerpViewingTimeLogs()
	if err != nil {
		return nil, err
	}

	b, err := l.store.ExportSerpViewingTimeLog(data, header, filetype)
	if err != nil {
		return nil, err
	}

	return b, nil

}

func (l *LogImpl) ExportPageViewingTimeLog(header bool, filetype store.FileType) (*bytes.Buffer, error) {
	if l == nil {
		return nil, errors.New("LogImpl is nil")
	}

	data, err := l.repository.FetchAllPageViewingTimeLogs()
	if err != nil {
		return nil, err
	}

	b, err := l.store.ExportPageViewingTimeLog(data, header, filetype)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (l *LogImpl) ExportSerpEventLog(header bool, filetype store.FileType) (*bytes.Buffer, error) {
	if l == nil {
		return nil, errors.New("LogImpl is nil")
	}

	data, err := l.repository.FetchAllSerpEventLogs()
	if err != nil {
		return nil, err
	}

	b, err := l.store.ExportSerpEventLog(data, header, filetype)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (l *LogImpl) ExportSearchSeeion(header bool, filetype store.FileType) (*bytes.Buffer, error) {
	if l == nil {
		return nil, errors.New("LogImpl is nil")
	}

	data, err := l.repository.FetchAllSearchSessions()
	if err != nil {
		return nil, err
	}

	b, err := l.store.ExportSearchSessionLog(data, header, filetype)
	if err != nil {
		return nil, err
	}

	return b, nil
}
