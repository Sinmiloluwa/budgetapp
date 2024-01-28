package main

import (
	"time"

	"github.com/Sinmiloluwa/budgetapp/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ApiKey    string `json:"token"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email: dbUser.Email,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}