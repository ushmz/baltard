//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"ratri/domain/model"
)

type SerpRepository interface {
	FetchSerpByTaskID(taskId, offset int) (*[]model.SearchPage, error)
}
