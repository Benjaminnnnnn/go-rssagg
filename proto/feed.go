package proto

import (
	"time"

	"github.com/benjaminnnnnn/go-review/rssagg/models"
	"github.com/google/uuid"
)

var ExampleCreateFeedResponse = models.Feed{
	ID:        uuid.New(),
	CreatedAt: time.Now().UTC(),
	UpdatedAt: time.Now().UTC(),
	Name:      "Some other rss feed title",
	Url:       "https://rss.nytimes.com/services/xml/rss/nyt/US.xml",
	UserID:    uuid.New(),
}
