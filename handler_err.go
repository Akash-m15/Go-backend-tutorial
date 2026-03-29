package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondErr(w, 200, "Your Mistake")
}
