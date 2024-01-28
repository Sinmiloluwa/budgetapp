package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sinmiloluwa/budgetapp/internal/auth"
	"github.com/Sinmiloluwa/budgetapp/internal/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:name`;
		Password string `json:password`;
		Email string `json:email`;
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	hashedPassword, err := HashPassword((params.Password))
	if err != nil {
		fmt.Println("Error hashing password:", err)
	}
	
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Password: hashedPassword,
		Email: params.Email,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create resource: %v", err))
	}
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get User: %v", err))
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}