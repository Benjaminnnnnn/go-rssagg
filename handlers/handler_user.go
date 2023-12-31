package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/benjaminnnnnn/go-review/rssagg/internal/database"
	"github.com/benjaminnnnnn/go-review/rssagg/models"
)

type ApiConfig struct {
	DB *database.Queries
}

func (apiCfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
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
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DBUserToUser(dbUser))
}

func (apiCfg *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJSON(w, http.StatusOK, models.DBUserToUser(user))
}

func (apiCfg *ApiConfig) GetPostsForUser(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	dbPosts,
		err := apiCfg.DB.GetPostsForUser(
		r.Context(),
		database.GetPostsForUserParams{UserID: user.ID, Limit: 10},
	)
	if err != nil {
		RespondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Unable to get posts for user %v", err),
		)
		return
	}

	RespondWithJSON(w, http.StatusOK, models.DBPostsToPosts(dbPosts))
}
