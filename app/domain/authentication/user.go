package authentication

import "time"

// UserAuthentication : Interface for the user authentication.
type UserAuthentication interface {
	// RegisterUser : RegisterUser user to authentication service.
	RegisterUser(externalID, secret string) error

	// Login(externalID, secret string) error

	// GenerateToken : Generate new token with externalID as an UID.
	GenerateToken(externalID string) (string, error)

	// GenerateSessionCookie : Generate new session cookie.
	GenerateSessionCookie(externalID string, expiresIn time.Duration) (string, error)

	// RevokeToken : Revoke generated token with externalID as an UID.
	RevokeToken(externalID string) error

	// DeleteUser : DeleteUser user from authentication service.
	DeleteUser(externalID string) error
}
