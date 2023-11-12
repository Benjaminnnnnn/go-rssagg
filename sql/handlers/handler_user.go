package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/benjaminnnnnn/go-review/rssagg/internal/auth"
	"github.com/benjaminnnnnn/go-review/rssagg/internal/database"
	"github.com/benjaminnnnnn/go-review/rssagg/models"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DB *database.Queries
}

type parameters struct {
	Name string `json:"name"`
}

func (apiCfg *ApiConfig) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	dbUser, err := apiCfg.DB.CreateUser(
		r.Context(), database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      params.Name,
		},
	)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %v", err))
	}

	RespondWithJSON(w, http.StatusCreated, models.DBUserToUser(dbUser))
}

func (apiCfg *ApiConfig) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, models.DBUserToUser(user))
}
