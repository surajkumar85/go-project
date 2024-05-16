package user

import (
	"encoding/json"
	"net/http"
)


type Handler struct {
}

func NewUserHandler() *Handler{
	return &Handler{}
}

func (h *Handler) RegisteredRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /",h.apiWelcomeRoute)
}

func (h *Handler) apiWelcomeRoute(w http.ResponseWriter,r *http.Request){

	 w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message":"this is the v1 of our golang api project"})
}