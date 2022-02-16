//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"ratri/internal/domain/model"
)

type LogRepository interface {
	// FetchAllSerpViewingTimeLogs : Fetch all `SerpViewingLog` data.
	// Please make sure that this method is used only for exporting data.
	FetchAllSerpViewingTimeLogs() ([]model.SerpViewingLog, error)

	// FetchAllPageViewingTimeLogs : Fetch all `PageViewingLog` data.
	// Please make sure that this method is used only for exporting data.
	FetchAllPageViewingTimeLogs() ([]model.PageViewingLog, error)

	// FetchAllSerpEventLogs : Fetch all `SerpEventLog` data.
	// Please make sure that this method is used only for exporting data.
	FetchAllSerpEventLogs() ([]model.SearchPageEventLog, error)

	// FetchAllSearchSessions : Fetch all `SearchSession` data.
	// Please make sure that this method is used only for exporting data.
	FetchAllSearchSessions() ([]model.SearchSession, error)

	// CumulateSerpViewingTime : Upsert serp viewing time log.
	CumulateSerpViewingTime(model.SerpViewingLogParam) error

	// CumulatePageViewingTime : Upsert page viewing time log.
	CumulatePageViewingTime(model.PageViewingLogParam) error

	// StoreSerpEventLog: Insert event log (such as click, paginate ...).
	StoreSerpEventLog(model.SearchPageEventLogParam) error

	// StoreSearchSeeion : Upsert searh session log.
	StoreSearchSeeion(model.SearchSessionParam) error
}
