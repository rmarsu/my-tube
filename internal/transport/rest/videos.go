package rest

import (
	"encoding/json"
	"myTube/internal/models"
	"myTube/internal/service"
	"myTube/pkg/log"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetTrendingVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := h.services.VideoService.GetTrendyVideos(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(videos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetVideoByID(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || videoId < 1 {
		log.Debug(err.Error())
		http.Error(w, "Invalid video ID", http.StatusBadRequest)
		return
	}
	video, err := h.services.VideoService.GetVideo(r.Context(), videoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Debug(video)
	json.NewEncoder(w).Encode(video)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateVideo(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Please login", http.StatusUnauthorized)
		return
	}
	id, err := h.services.Users.GetUserIdFromToken(r.Context(), authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var videoSettings service.VideoInput
	err = json.NewDecoder(r.Body).Decode(&videoSettings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	thumbnailFilePath := ""
	videoFilePath := ""
	// TODO upload thumbnail file logic here
	// TODO add file upload logic here

	video := models.Video{
		AuthorID:    userId,
		Title:       videoSettings.Title,
		Description: videoSettings.Description,
		Filepath:    videoFilePath,
		Thumbnail:   thumbnailFilePath,
		Views:       0,
		Likes:       0,
		CreatedAt:   time.Now(),
	}

	err = h.services.VideoService.CreateVideo(r.Context(), video)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(video)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateVideo(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Please login", http.StatusUnauthorized)
		return
	}
	id, err := h.services.Users.GetUserIdFromToken(r.Context(), authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	videoId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || videoId < 0 {
		http.Error(w, "Invalid video ID", http.StatusBadRequest)
		return
	}
	video, err := h.services.VideoService.GetVideo(r.Context(), videoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if video.AuthorID != userId {
		http.Error(w, "Unauthorized to update this video", http.StatusUnauthorized)
		return
	}
	var videoSettings service.VideoInput
	err = json.NewDecoder(r.Body).Decode(&videoSettings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	video.Title = videoSettings.Title
	video.Description = videoSettings.Description
	err = h.services.VideoService.UpdateVideo(r.Context(), videoId, video)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(video)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Please login", http.StatusUnauthorized)
		return
	}
	id, err := h.services.Users.GetUserIdFromToken(r.Context(), authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	videoId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || videoId < 0 {
		http.Error(w, "Invalid video ID", http.StatusBadRequest)
		return
	}

	video, err := h.services.VideoService.GetVideo(r.Context(), videoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if video.AuthorID != userId {
		http.Error(w, "Unauthorized to delete this video", http.StatusUnauthorized)
		return
	}
	err = h.services.VideoService.DeleteVideo(r.Context(), videoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) SearchVideos(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}

	videos, err := h.services.VideoService.SearchVideos(r.Context(), query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(videos)
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) LikeVideo(w http.ResponseWriter, r *http.Request) {
     authHeader := r.Header.Get("Authorization")
     if authHeader == "" {
         http.Error(w, "Please login", http.StatusUnauthorized)
         return
     }
     id, err := h.services.Users.GetUserIdFromToken(r.Context(), authHeader)
     if err!= nil {
         http.Error(w, err.Error(), http.StatusUnauthorized)
         return
     }
     userId, err := strconv.Atoi(id)
     if err!= nil {
         http.Error(w, "Invalid user ID", http.StatusBadRequest)
         return
     }
     videoId, err := strconv.Atoi(r.URL.Query().Get("id"))
     if err!= nil || videoId < 0 {
         http.Error(w, "Invalid video ID", http.StatusBadRequest)
         return
     }
     err = h.services.VideoService.LikeVideo(r.Context(), videoId, userId)
     if err!= nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
     }
}
