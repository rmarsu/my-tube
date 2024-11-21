package rest

import (
	"encoding/json"
	"fmt"
	"myTube/internal/service"
	"net/http"
)


func (h *Handler) GetVideo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
     fmt.Fprintf(w, "Get Video Handler")
}

func (h *Handler) CreateVideo(w http.ResponseWriter, r *http.Request) {
     w.WriteHeader(http.StatusCreated)
     fmt.Fprintf(w, "Create Video Handler")
}

func (h *Handler) DeleteVideo(w http.ResponseWriter, r *http.Request) {
     w.WriteHeader(http.StatusNoContent)
     fmt.Fprintf(w, "Delete Video Handler")
}


func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
     var input service.UserSignInInput

     err := json.NewDecoder(r.Body).Decode(&input)
     if err != nil {
          http.Error(w, err.Error(), http.StatusBadRequest)
          return
     }
     tokens, err := h.services.Users.SignIn(r.Context(), service.UserSignInInput{
          Username: input.Username,
          Password: input.Password,
     })
     if err!= nil {
          http.Error(w, err.Error(), http.StatusUnauthorized)
          return
     }
     json.NewEncoder(w).Encode(tokens)
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
     var input service.UserSignUpInput
     err := json.NewDecoder(r.Body).Decode(&input)
     if err!= nil {
          http.Error(w, err.Error(), http.StatusBadRequest)
          return
     }
     tokens , err := h.services.Users.SignUp(r.Context(), input)
     if err!= nil {
          http.Error(w, err.Error(), http.StatusUnauthorized)
          return
     }
     json.NewEncoder(w).Encode(tokens)

}
