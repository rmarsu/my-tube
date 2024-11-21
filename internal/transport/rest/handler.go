package rest

import (
	"myTube/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
     return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/videos", h.GetVideo)
	mux.HandleFunc("/videos/{id}", h.GetVideo)
	mux.HandleFunc("/create/video", h.CreateVideo)
	
	mux.HandleFunc("/signup", h.SignUp)
	mux.HandleFunc("/signin", h.SignIn)
	return mux
}
