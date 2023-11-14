package models

import (
	"time"

	"github.com/benjaminnnnnn/go-review/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" example:"d7148b26-9a35-4e8f-9a84-71ebf7661b38"`
	CreatedAt time.Time `json:"created_at" example:"2009-11-10 23:00:00 +0000 UTC"`
	UpdatedAt time.Time `json:"updated_at" example:"2009-11-10 23:00:00 +0000 UTC"`
	Name      string    `json:"name" example:"John Seed"`
	ApiKey    string    `json:"api_key" example:"1c2071838a1d50760470ed7fbcd30dd1a53bea0cd875ca27f7e3c58447970c2c"`
}

func DBUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}
