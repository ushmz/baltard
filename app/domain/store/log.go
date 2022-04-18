package store

import (
	"bytes"
	"ratri/domain/model"
)

// FileType : Export file type ("csv", "tsv")
type FileType string

const (
	// CSV : separated by ","
	CSV = FileType("csv")
	// TSV : separated by "\t"
	TSV = FileType("tsv")
)

// LogStore : Abstract operations that `Log` model should have.
type LogStore interface {
	ExportSerpDwellTimeLog([]model.SerpDwellTimeLog, bool, FileType) (*bytes.Buffer, error)
	ExportPageDwellTimeLog([]model.PageDwellTimeLog, bool, FileType) (*bytes.Buffer, error)
	ExportSerpEventLog([]model.SearchPageEventLog, bool, FileType) (*bytes.Buffer, error)
	ExportSearchSessionLog([]model.SearchSession, bool, FileType) (*bytes.Buffer, error)
}
