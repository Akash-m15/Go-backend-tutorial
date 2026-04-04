package main

import (
	"fmt"
	"net/http"

	"github.com/Akash-m15/rssagg/internal/auth"
	"github.com/Akash-m15/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondErr(w, 403, fmt.Sprintf("Error is getting APIKey: %v", err))
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), api_key)
		if err != nil {
			respondErr(w, 400, fmt.Sprintf("Error in getting User : %v", err))
		}
		handler(w, r, user)
	}
}
