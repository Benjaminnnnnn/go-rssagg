package handlers

import "net/http"

func HandlerErr(w http.ResponseWriter, req *http.Request) {
	RespondWithError(w, 400, "Something went wrong")
}
