package fileio

import (
	"bytes"
	"encoding/csv"
	"ratri/internal/domain/model"
	"ratri/internal/domain/store"
	"strconv"
)

type LogStore struct{}

func NewLogStore() *LogStore {
	return &LogStore{}
}

func (l *LogStore) ExportSerpViewingTimeLog(data []model.SerpViewingLog, header bool, filetype store.FileType) ([]byte, error) {
	content := [][]string{}
	if header {
		content = append(content, []string{
			"user_id",
			"task_id",
			"condition_id",
			"dwell_time",
			"created_at",
			"updated_at",
		})
	}
	for _, v := range data {
		content = append(content, []string{
			strconv.Itoa(v.UserID),
			strconv.Itoa(v.TaskID),
			strconv.Itoa(v.ConditionID),
			strconv.Itoa(v.DwellTime),
			v.CreatedAt.Format("2006-01-02 13:34:56"),
			v.UpdatedAt.Format("2006-01-02 13:34:56"),
		})
	}
	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	if filetype == store.TSV {
		w.Comma = '\t'
	}
	err := w.WriteAll(content)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (l *LogStore) ExportPageViewingTimeLog(data []model.PageViewingLog, header bool, filetype store.FileType) ([]byte, error) {
	content := [][]string{}
	if header {
		content = append(content, []string{
			"user_id",
			"task_id",
			"page_id",
			"condition_id",
			"dwell_time",
			"created_at",
			"updated_at",
		})
	}
	for _, v := range data {
		content = append(content, []string{
			strconv.Itoa(v.UserID),
			strconv.Itoa(v.TaskID),
			strconv.Itoa(v.PageID),
			strconv.Itoa(v.ConditionID),
			strconv.Itoa(v.DwellTime),
			v.CreatedAt.Format("2006-01-02 13:34:56"),
			v.UpdatedAt.Format("2006-01-02 13:34:56"),
		})
	}
	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	if filetype == store.TSV {
		w.Comma = '\t'
	}
	err := w.WriteAll(content)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (l *LogStore) ExportSerpEventLog(data []model.SearchPageEventLog, header bool, filetype store.FileType) ([]byte, error) {
	content := [][]string{}
	if header {
		content = append(content, []string{
			"user_id",
			"task_id",
			"page_id",
			"condition_id",
			"dwell_time",
			"created_at",
			"updated_at",
		})
	}

	for _, v := range data {
		content = append(content, []string{
			v.ID,
			strconv.Itoa(v.UserID),
			strconv.Itoa(v.TaskID),
			strconv.Itoa(v.ConditionID),
			strconv.Itoa(v.Time),
			strconv.Itoa(v.Page),
			strconv.Itoa(v.Rank),
			strconv.FormatBool(v.IsVisible),
			v.Event,
			v.CreatedAt.Format("2006-01-02 13:34:56"),
			v.UpdatedAt.Format("2006-01-02 13:34:56"),
		})
	}
	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	if filetype == store.TSV {
		w.Comma = '\t'
	}
	err := w.WriteAll(content)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (l *LogStore) ExportSearchSessionLog(data []model.SearchSession, header bool, filetype store.FileType) ([]byte, error) {
	content := [][]string{}
	if header {
		content = append(content, []string{
			"user_id",
			"task_id",
			"page_id",
			"condition_id",
			"dwell_time",
			"created_at",
			"updated_at",
		})
	}

	for _, v := range data {
		content = append(content, []string{
			strconv.Itoa(v.UserID),
			strconv.Itoa(v.TaskID),
			strconv.Itoa(v.ConditionID),
			v.StartedAt.Format("2006-01-02 13:34:56"),
			v.EndedAt.Format("2006-01-02 13:34:56"),
		})
	}
	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	if filetype == store.TSV {
		w.Comma = '\t'
	}
	err := w.WriteAll(content)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
