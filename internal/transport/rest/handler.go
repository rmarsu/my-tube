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
	mux.HandleFunc("/", h.GetTrendingVideos)
	mux.HandleFunc("/videos", h.GetVideoByID)
	mux.HandleFunc("/upload/video", h.CreateVideo)
	mux.HandleFunc("/update/video", h.UpdateVideo)
	mux.HandleFunc("/delete/video", h.DeleteVideo)
	mux.HandleFunc("/search", h.SearchVideos)
	mux.HandleFunc("/like", h.LikeVideo)
	
	mux.HandleFunc("/signup", h.SignUp)
	mux.HandleFunc("/signin", h.SignIn)
	return mux
}
