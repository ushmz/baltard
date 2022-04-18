package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
)

// CookieAuth : Auth middleware that check cookie and verify token
func CookieAuth(app *firebase.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			if ctx == nil {
				ctx = context.Background()
			}

			client, err := app.Auth(ctx)
			if err != nil {
				fmt.Printf("Failed to get auth client : %v\n", err)
				return errors.New("Failed to get auth client")
			}

			cookie, err := c.Request().Cookie("exp-session")
			if err != nil {
				if err == http.ErrNoCookie {
					fmt.Printf("Failed to get cookie value (No cookie) : %v\n", err)
					return errors.New("No cookie")
				}
				fmt.Printf("Failed to get cookie value : %v\n", err)
				return errors.New("Invalid request")
			}

			token, err := client.VerifySessionCookieAndCheckRevoked(ctx, cookie.Value)
			if err != nil {
				fmt.Printf("error verifying ID token: %v\n", err)
				return errors.New("Invalid token")
			}

			c.Set("token", token)
			return next(c)
		}
	}
}
