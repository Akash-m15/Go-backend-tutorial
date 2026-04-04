package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Akash-m15/rssagg/internal/database"
	"github.com/go-chi/chi"
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
	userRes := dbUserToUser(user)
	respondWithJSON(w, 200, userRes)
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	userRes := dbUserToUser(user)
	respondWithJSON(w, 200, userRes)
}

func (apiCfg *apiConfig) handlerGetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := chi.URLParam(r, "limit")
	limit, err := strconv.ParseInt(limitStr, 0, 32)
	if err != nil {
		respondErr(w, 403, "Error parsing limit")
	}

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondErr(w, 400, fmt.Sprintf("Error in creating User : %v", err))
	}
	respondWithJSON(w, 200, dbPostsToPosts(posts))
}
