package firebase

import (
	"context"
	"ratri/domain/authentication"
	"ratri/domain/model"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/xerrors"
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
		return xerrors.Errorf("UserAuthenticationImpl.RegisterUser() is called with nil receiver: %w", model.ErrNilReceiver)
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return xerrors.Errorf("Failed to get auth client: %w", err)
	}

	params := (&auth.UserToCreate{}).
		UID(externalID).
		Email(externalID + "@savitr.dummy.com").
		EmailVerified(true).
		Password(secret)

	if _, err = client.CreateUser(ctx, params); err != nil {
		if auth.IsUIDAlreadyExists(err) {
			return xerrors.Errorf("Given ID is already used: %w", err)
		}
		return xerrors.Errorf("Failed to create user: %w", err)
	}

	return nil
}

// DeleteUser : Delete user from application.
func (u *UserAuthenticationImpl) DeleteUser(externalID string) error {
	if u == nil {
		return xerrors.Errorf("UserAuthenticationImpl.DeleteUser() is called with nil receiver: %w", model.ErrNilReceiver)
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return xerrors.Errorf("Failed to get auth client: %w", err)
	}

	if err := client.DeleteUser(context.Background(), externalID); err != nil {
		return xerrors.Errorf("Failed to delete user: %w", err)
	}
	return nil
}

// GenerateToken : Generate new token with externalID as an UID.
func (u *UserAuthenticationImpl) GenerateToken(externalID string) (string, error) {
	if u == nil {
		return "", xerrors.Errorf("UserAuthenticationImpl.GenerateToken() is called with nil receiver: %w", model.ErrNilReceiver)
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return "", xerrors.Errorf("Failed to get auth client: %w", err)
	}

	token, err := client.CustomToken(ctx, externalID)
	if err != nil {
		return "", xerrors.Errorf("Failed to generate user token: %w", err)
	}
	return token, nil
}

// GenerateSessionCookie : Generate new session cookie with externalID as an UID.
func (u *UserAuthenticationImpl) GenerateSessionCookie(idToken string, expiresIn time.Duration) (string, error) {
	if u == nil {
		return "", xerrors.Errorf("UserAuthenticationImpl.GenerateSessionCookie() is called with nil receiver: %w", model.ErrNilReceiver)
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return "", xerrors.Errorf("Failed to get auth client: %w", err)
	}

	cookie, err := client.SessionCookie(ctx, idToken, expiresIn)
	if err != nil {
		return "", xerrors.Errorf("Failed to generate auth cookie: %w", err)
	}
	return cookie, nil
}

// RevokeToken : Revoke generated token with externalID as an UID.
func (u *UserAuthenticationImpl) RevokeToken(externalID string) error {
	if u == nil {
		return xerrors.Errorf("UserAuthenticationImpl.RevokeToken() is called with nil receiver: %w", model.ErrNilReceiver)
	}

	ctx := context.Background()
	client, err := u.App.Auth(ctx)
	if err != nil {
		return xerrors.Errorf("Failed to get auth client: %w", err)
	}

	if err := client.RevokeRefreshTokens(ctx, externalID); err != nil {
		return xerrors.Errorf("Failed to revoke token: %w", err)
	}
	return nil
}
