package middlewares

import (
	"fmt"
	"net/http"

	"github.com/benjaminnnnnn/go-review/rssagg/handlers"
	"github.com/benjaminnnnnn/go-review/rssagg/internal/auth"
	"github.com/benjaminnnnnn/go-review/rssagg/internal/database"
)

type AuthMiddleware struct {
	ApiConfig *handlers.ApiConfig
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (m *AuthMiddleware) MiddlewareAuth(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			handlers.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := m.ApiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			handlers.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		next(w, r, user)
	}
}
