//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"ratri/internal/domain/model"
)

type UserRepository interface {
	Create(uid, secret string) (*model.User, error)
	FindById(UserId int) (*model.User, error)
	FindByUid(uid string) (*model.User, error)
	AddCompletionCode(userId, code int) error
	GetCompletionCodeById(userId int) (int, error)
}
