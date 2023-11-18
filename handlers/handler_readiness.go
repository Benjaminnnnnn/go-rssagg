package handlers

import "net/http"

type EmptyResponse struct{}

// System Health godoc
// @Summary      Display an error message
// @Accept       json
// @Produce      json
// @Success      200  {object}	EmptyResponse "Server ready"
// @Router       /healthz [get]
func HandlerReadiness(w http.ResponseWriter, req *http.Request) {
	RespondWithJSON(w, http.StatusOK, EmptyResponse{})
}
