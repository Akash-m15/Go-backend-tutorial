package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct {
		Message string
		Code    string
	}{
		Message: "Message",
		Code:    "1111",
	})
}
