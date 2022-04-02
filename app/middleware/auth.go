package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
)

func Auth(app *firebase.App) echo.MiddlewareFunc {
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

			auth := c.Request().Header.Get("Authorization")
			idToken := strings.Replace(auth, "Bearer ", "", 1)

			token, err := client.VerifyIDToken(ctx, idToken)
			if err != nil {
				fmt.Printf("error verifying ID token: %v\n", err)
				return errors.New("Invalid token")
			}

			c.Set("token", token)
			return next(c)
		}
	}
}
