package main

import (
	"github.com/Akash-m15/rssagg/internal/database"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	ApiKey string    `json:"api_key"`
}

func dbUserToUser(dbUser database.User) UserResponse {
	return UserResponse{
		ID:     dbUser.ID,
		Name:   dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}

type feedResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Url    string    `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

func dbFeedToFeed(dbFeed database.Feed) feedResponse {
	return feedResponse{
		ID:     dbFeed.ID,
		Name:   dbFeed.Name,
		Url:    dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func dbFeedsToFeeds(dbFeeds []database.Feed) []feedResponse {
	feeds := []feedResponse{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, dbFeedToFeed(feed))
	}
	return feeds
}

type feedFollowResponse struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}

func dbFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) feedFollowResponse {
	return feedFollowResponse{
		ID:     dbFeedFollow.ID,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}
}

func dbFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []feedFollowResponse {
	feedFollows := []feedFollowResponse{}
	for _, feedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, dbFeedFollowToFeedFollow(feedFollow))
	}
	return feedFollows
}

type PostResponse struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Url          string    `json:"url"`
	Published_At string    `json:"published_at"`
}

func dbPostToPost(dbPost database.Post) PostResponse {
	return PostResponse{
		ID:           dbPost.ID,
		Title:        dbPost.Title,
		Description:  dbPost.Description.String,
		Url:          dbPost.Url,
		Published_At: dbPost.PublishedAt.String(),
	}
}

func dbPostsToPosts(dbPosts []database.Post) []PostResponse {
	posts := []PostResponse{}

	for _, post := range dbPosts {
		posts = append(posts, dbPostToPost(post))
	}
	return posts
}
