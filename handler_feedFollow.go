package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Akash-m15/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondErr(w, 400, "Invalid postBody")
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondErr(w, 400, fmt.Sprintf("Error in creating FeedFollow : %v", err))
	}
	respondWithJSON(w, 200, dbFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondErr(w, 404, fmt.Sprintf("Error in getting feedFollow: %v", err))
	}

	respondWithJSON(w, 200, dbFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowId")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondErr(w, 404, fmt.Sprintf("Error in parsing feedFollow Id: %v", err))
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondErr(w, 404, fmt.Sprintf("Error in deleting feedFollow: %v", err))
	}

	respondWithJSON(w, 200, struct {
		Msg string `json:"message"`
	}{
		Msg: "Deleted Feed Follow",
	})
}
