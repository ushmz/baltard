package store

import "ratri/internal/domain/model"

type FileType string

const (
	CSV = FileType("csv")
	TSV = FileType("tsv")
)

type LogStore interface {
	ExportSerpViewingTimeLog([]model.SerpViewingLog, bool, FileType) ([]byte, error)
	ExportPageViewingTimeLog([]model.PageViewingLog, bool, FileType) ([]byte, error)
	ExportSerpEventLog([]model.SearchPageEventLog, bool, FileType) ([]byte, error)
	ExportSearchSessionLog([]model.SearchSession, bool, FileType) ([]byte, error)
}
