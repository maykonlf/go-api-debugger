package handlers

import "net/http"

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methos", "GET, POST, PUT, DELETE, OPTIONS")

	w.WriteHeader(http.StatusOK)
}