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

// @Summary		Create an RSS feed
// @Description	Create an RSS feed provided by user
// @Tags		feeds
// @Accept		json
// @Produce		json
// @Param		Authorization	header	string 		true	"ApiKey {api_key}"
// @Param		body	body	handlers.CreateFeed.parameters	true	"RSS Feed body"
// @Success		201	{object}	models.Feed{proto.ExampleCreateFeedResponse}
// @Failure		400	{object}	HTTPError		"Incorrect reuqest body"
// @Failure		400	{object}	HTTPError		"Couldn't create feed"
// @Security	ApiKeyAuth
// @Router		/feeds	[post]
func (apiCfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name" example:"New York Times - U.S."`
		Url  string `json:"url" example:"https://rss.nytimes.com/services/xml/rss/nyt/US.xml"`
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

// @Summary		List all RSS feeds
// @Description	List all RSS feeds
// @Tags		feeds
// @Produce		json
// @Success		201	{array}		models.Feed
// @Failure		400	{object}	HTTPError		"Couldn't get feeds"
// @Router		/feeds	[get]
func (apiCfg *ApiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	dbFeeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feeds."))
		return
	}

	feeds := models.DBFeedsToFeeds(dbFeeds)

	RespondWithJSON(w, http.StatusOK, feeds)
}
