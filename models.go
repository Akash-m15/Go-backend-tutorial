package main

import (
	"github.com/Akash-m15/rssagg/internal/database"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func dbUserToUser(dbUser database.User) UserResponse {
	return UserResponse{
		ID:   dbUser.ID,
		Name: dbUser.Name,
	}
}
