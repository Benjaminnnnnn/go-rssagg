package handlers

import "net/http"

func HandlerReadiness(w http.ResponseWriter, req *http.Request) {
	RespondWithJSON(w, 200, struct{}{})
}
