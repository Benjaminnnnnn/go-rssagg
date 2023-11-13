package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/benjaminnnnnn/go-review/rssagg/internal/database"
	"github.com/benjaminnnnnn/go-review/rssagg/models"
	"github.com/google/uuid"
)

type feedParameters struct {
	Name string `json:"name"`
}

func (apiCfg *ApiConfig) HandleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := feedParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	dbFeed, err := apiCfg.DB.CreateFeed(
		r.Context(), database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      params.Name,
			Url:       "",
			UserID:    user.ID,
		},
	)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %v", err))
	}

	RespondWithJSON(w, http.StatusCreated, models.DBFeddToFeed(dbFeed))
}
