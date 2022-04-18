//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import "ratri/domain/model"

// LinkedPageRepository : Abstract operations that `LinkedPage` model should have.
type LinkedPageRepository interface {
	Get(linkedPageID int) (model.LinkedPage, error)
	GetBySearchPageIDs(pageID []int, taskID, top int) (*[]model.SearchPageWithLinkedPage, error)
	// [TODO] Better name
	GetRatioBySearchPageIDs(pageID []int, taskID int) (*[]model.SearchPageWithLinkedPageRatio, error)
	Select(linkedPageIDs []int) (*[]model.LinkedPage, error)
	List(offset, limit int) (*[]model.LinkedPage, error)
	// Following methods are not implemented
	// because this `LinkedPage` resource should not be created by API.
	// Create(model.LinkedPage) (model.LinkedPage, error)
	// Update(model.LinkedPage) (model.LinkedPage, error)
	// Delete(linkedPageID int) error
}
