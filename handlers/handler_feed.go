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

func (apiCfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	dbFeed, err := apiCfg.DB.CreateFeed(
		r.Context(), database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      params.Name,
			Url:       params.Url,
			UserID:    user.ID,
		},
	)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DBFeedToFeed(dbFeed))
}

func (apiCfg *ApiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	dbFeeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feeds."))
		return
	}

	feeds := models.DBFeedsToFeeds(dbFeeds)

	RespondWithJSON(w, http.StatusOK, feeds)
}
