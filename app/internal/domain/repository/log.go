//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"ratri/internal/domain/model"
)

type LogRepository interface {
	StoreTaskTimeLog(*model.SerpViewingLogParamWithTime) error
	CumulateSerpViewingTime(*model.SerpViewingLogParam) error
	CumulatePageViewingTime(*model.PageViewingLogParam) error
	StoreSerpEventLog(*model.SearchPageEventLogParam) error
	StoreSearchSeeion(*model.SearchSession) error
}
