package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rohan3011/go-server/utils"
)

type contextKey string

const UserKey = contextKey("user")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session token from the cookie
		cookie, err := r.Cookie("session_token")
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			return
		}

		// Validate the token
		user, err := ValidateSessionToken(cookie.Value)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			return
		}

		// Set the user in the request context for further use
		ctx := context.WithValue(r.Context(), UserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
