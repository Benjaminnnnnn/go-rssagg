package middlwares

import (
	"net/http"

	"github.com/benjaminnnnnn/go-review/rssagg/internal/database"
	"github.com/benjaminnnnnn/go-review/rssagg/sql/handlers"
)

type AuthMiddleware struct {
	ApiConfig *handlers.ApiConfig
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (m *AuthMiddleware) middlewareAuth(handler authedHandler) http.HandlerFunc
