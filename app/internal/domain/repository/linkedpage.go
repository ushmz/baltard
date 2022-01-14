//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import "ratri/internal/domain/model"

type LinkedPageRepository interface {
	Get(linkedPageId int) (model.LinkedPage, error)
	GetBySearchPageId(pageId, taskId, top int) (*[]model.LinkedPage, error)
	Select(linkedPageIds []int) (*[]model.LinkedPage, error)
	List(offset, limit int) (*[]model.LinkedPage, error)
	// Following methods are not implemented
	// because this `LinkedPage` resource should not be created by API.
	// Create(model.LinkedPage) (model.LinkedPage, error)
	// Update(model.LinkedPage) (model.LinkedPage, error)
	// Delete(linkedPageId int) error
}
