package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Akash-m15/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondErr(w, 400, "Invalid postBody")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondErr(w, 400, fmt.Sprintf("Error in creating feed : %v", err))
	}
	feedRes := dbFeedToFeed(feed)
	respondWithJSON(w, 200, feedRes)
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context(), user.ID)

	if err != nil {
		respondErr(w, 403, fmt.Sprintf("Error is getting feeds: %v", err))
	}
	respondWithJSON(w, 200, dbFeedsToFeeds(feeds))
}
