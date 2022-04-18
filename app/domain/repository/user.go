//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"ratri/domain/model"
)

// UserRepository : Abstract operations that `User` model should have.
type UserRepository interface {
	Create(uid string) (model.User, error)
	FindByID(UserID int) (model.User, error)
	FindByUID(uid string) (model.User, error)
	AddCompletionCode(userID, code int) error
	GetCompletionCodeByID(userID int) (int, error)
}
