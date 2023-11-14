package models

import (
	"time"

	"github.com/benjaminnnnnn/go-review/rssagg/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id" exmaple:"d7148b26-9a35-4e8f-9a84-71ebf7661b38"`
	CreatedAt time.Time `json:"created_at" example:"2009-11-10 23:00:00 +0000 UTC"`
	UpdatedAt time.Time `json:"updated_at" example:"2009-11-10 23:00:00 +0000 UTC"`
	Name      string    `json:"name" example:"New York Times"`
	Url       string    `json:"url" example:"https://rss.nytimes.com/services/xml/rss/nyt/US.xml"`
	UserID    uuid.UUID `json:"user_id" example:"54bd9e94-107c-40a4-91d0-973e0cef143d"`
}

func DBFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func DBFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, DBFeedToFeed(dbFeed))
	}
	return feeds
}
