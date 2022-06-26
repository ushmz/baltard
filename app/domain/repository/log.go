//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"ratri/domain/model"
)

// LogRepository : Abstract operations that `Log` model should have.
type LogRepository interface {
	// FetchAllSerpDwellTimeLogs : Fetch all `SerpDwellLog` data.
	// Please make sure that this method is used only for exporting data.
	FetchAllSerpDwellTimeLogs() ([]model.SerpDwellTimeLog, error)

	// FetchAllPageDwellTimeLogs : Fetch all `PageDwellLog` data.
	// Please make sure that this method is used only for exporting data.
	FetchAllPageDwellTimeLogs() ([]model.PageDwellTimeLog, error)

	// FetchAllSerpEventLogs : Fetch all `SerpEventLog` data.
	// Please make sure that this method is used only for exporting data.
	FetchAllSerpEventLogs() ([]model.SearchPageEventLog, error)

	// FetchAllSearchSessions : Fetch all `SearchSession` data.
	// Please make sure that this method is used only for exporting data.
	FetchAllSearchSessions() ([]model.SearchSession, error)

	// CumulateSerpDwellTime : Upsert serp viewing time log.
	CumulateSerpDwellTime(*model.SerpDwellTimeLogParam) error

	// CumulatePageDwellTime : Upsert page viewing time log.
	CumulatePageDwellTime(*model.PageDwellTimeLogParam) error

	// StoreSerpEventLog: Insert event log (such as click, paginate ...).
	StoreSerpEventLog(*model.SearchPageEventLogParam) error

	// StoreSearchSeeion : Upsert searh session log.
	StoreSearchSeeion(*model.SearchSessionParam) error
}
