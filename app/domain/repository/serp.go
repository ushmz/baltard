//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"ratri/domain/model"
)

// SerpRepository : Abstract operations that `Serp` model should have.
type SerpRepository interface {
	FetchSerpByTaskID(taskID, offset int) ([]model.SearchPage, error)
}
