package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Akash-m15/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondErr(w, 400, "Invalid postBody")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondErr(w, 400, fmt.Sprintf("Error in creating User : %v", err))
	}
	fmt.Println(user)
	respondWithJSON(w, 200, user)
}
