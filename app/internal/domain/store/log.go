package store

import (
	"bytes"
	"ratri/internal/domain/model"
)

type FileType string

const (
	CSV = FileType("csv")
	TSV = FileType("tsv")
)

type LogStore interface {
	ExportSerpViewingTimeLog([]model.SerpViewingLog, bool, FileType) (*bytes.Buffer, error)
	ExportPageViewingTimeLog([]model.PageViewingLog, bool, FileType) (*bytes.Buffer, error)
	ExportSerpEventLog([]model.SearchPageEventLog, bool, FileType) (*bytes.Buffer, error)
	ExportSearchSessionLog([]model.SearchSession, bool, FileType) (*bytes.Buffer, error)
}
