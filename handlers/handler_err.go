package handlers

import (
	"net/http"
)

// Test Error godoc
// @Summary      Display an error message
// @Accept       json
// @Produce      json
// @Failure      400  {object} HTTPError "Something went wrong"
// @Router       /error [get]
func HandlerErr(w http.ResponseWriter, req *http.Request) {
	RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
