package handler

import (
	"encoding/json"
	models "kilosale_main/internal/models/user"
	service "kilosale_main/internal/services/user"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.userService.Create(&user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func (h *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.userService.GetByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
