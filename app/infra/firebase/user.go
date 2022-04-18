package firebase

import (
	"context"
	"ratri/domain/authentication"
	"ratri/domain/model"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
)

// UserAuthenticationImpl : Implemention of user authentication.
type UserAuthenticationImpl struct {
	App *firebase.App
}

// NewUserAuthenticationImpl : Return new UserAuthentication implemention
func NewUserAuthenticationImpl(app *firebase.App) authentication.UserAuthentication {
	return &UserAuthenticationImpl{App: app}
}

// RegisterUser : Register user with externalID and password.
func (u *UserAuthenticationImpl) RegisterUser(externalID, secret string) error {
	if u == nil {
		return model.ErrNilReceiver
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return errors.New("Failed to get auth client")
	}

	params := (&auth.UserToCreate{}).
		UID(externalID).
		Email(externalID + "@savitr.dummy.com").
		EmailVerified(true).
		Password(secret)

	if _, err = client.CreateUser(ctx, params); err != nil {
		if auth.IsUIDAlreadyExists(err) {
			return errors.New("Given ID is already used")
		}
		return errors.New("Failed to create user")
	}

	return nil
}

// DeleteUser : Delete user from application.
func (u *UserAuthenticationImpl) DeleteUser(externalID string) error {
	if u == nil {
		return errors.WithStack(model.ErrNilReceiver)
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return errors.WithStack(errors.New("Failed to get auth client"))
	}

	if err := client.DeleteUser(context.Background(), externalID); err != nil {
		return errors.WithStack(errors.New("Failed to delete user"))
	}
	return nil
}

// GenerateToken : Generate new token with externalID as an UID.
func (u *UserAuthenticationImpl) GenerateToken(externalID string) (string, error) {
	if u == nil {
		return "", errors.WithStack(model.ErrNilReceiver)
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return "", errors.WithStack(errors.New("Failed to get auth client"))
	}

	token, err := client.CustomToken(ctx, externalID)
	if err != nil {
		return "", errors.WithStack(errors.New("Failed to generate user token"))
	}
	return token, nil
}

// GenerateSessionCookie : Generate new session cookie with externalID as an UID.
func (u *UserAuthenticationImpl) GenerateSessionCookie(idToken string, expiresIn time.Duration) (string, error) {
	if u == nil {
		return "", errors.WithStack(model.ErrNilReceiver)
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return "", errors.WithStack(err)
	}

	cookie, err := client.SessionCookie(ctx, idToken, expiresIn)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return cookie, nil
}

// RevokeToken : Revoke generated token with externalID as an UID.
func (u *UserAuthenticationImpl) RevokeToken(externalID string) error {
	if u == nil {
		return errors.WithStack(model.ErrNilReceiver)
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return errors.WithStack(errors.New("Failed to get auth client"))
	}

	if err := client.RevokeRefreshTokens(ctx, externalID); err != nil {
		return errors.WithStack(errors.New("Failed to revoke token"))
	}
	return nil
}
