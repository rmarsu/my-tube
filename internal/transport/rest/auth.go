package rest

import (
	"encoding/json"
	"myTube/internal/service"
	"net/http"
)

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
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input service.UserSignUpInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tokens, err := h.services.Users.SignUp(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(tokens)
}
