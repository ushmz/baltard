package fileio

import (
	"bytes"
	"encoding/csv"
	"ratri/domain/model"
	"ratri/domain/store"
	"strconv"
)

// LogStore : Implemention of log exporting
type LogStore struct{}

// NewLogStore : Return new log exporting struct
func NewLogStore() *LogStore {
	return &LogStore{}
}

// ExportSerpDwellTimeLog : Write all SERP dwell time log to buffer
func (l *LogStore) ExportSerpDwellTimeLog(data []model.SerpDwellTimeLog, header bool, filetype store.FileType) (*bytes.Buffer, error) {
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
			v.Created.Format("2006-01-02 13:34:56"),
			v.Updated.Format("2006-01-02 13:34:56"),
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
	return b, nil
}

// ExportPageDwellTimeLog : Write all result page dwell time log to buffer
func (l *LogStore) ExportPageDwellTimeLog(data []model.PageDwellTimeLog, header bool, filetype store.FileType) (*bytes.Buffer, error) {
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
			v.Created.Format("2006-01-02 13:34:56"),
			v.Updated.Format("2006-01-02 13:34:56"),
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
	return b, nil
}

// ExportSerpEventLog : Write all SERP event log to file
func (l *LogStore) ExportSerpEventLog(data []model.SearchPageEventLog, header bool, filetype store.FileType) (*bytes.Buffer, error) {
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
			v.Created.Format("2006-01-02 13:34:56"),
			v.Updated.Format("2006-01-02 13:34:56"),
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
	return b, nil
}

// ExportSearchSessionLog : Write all search session logs to buffer
func (l *LogStore) ExportSearchSessionLog(data []model.SearchSession, header bool, filetype store.FileType) (*bytes.Buffer, error) {
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
			v.Started.Format("2006-01-02 13:34:56"),
			v.Ended.Format("2006-01-02 13:34:56"),
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
	return b, nil
}
