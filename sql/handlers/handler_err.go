package handlers

import "net/http"

func HandlerErr(w http.ResponseWriter, req *http.Request) {
	RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
