package handler

import (
	"encoding/json"
	models "kilosale_main/internal/models/user"
	service "kilosale_main/internal/services/user"
	"net/http"

	"github.com/gechderib/kilosale/pkg/response"
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
		response.Success(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	err = h.userService.Create(&user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create user", nil)
		return
	}

	response.Success(w, http.StatusCreated, "User created successfully", user)

}

func (h *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.userService.GetByID(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve user", nil)
		return
	}

	response.Success(w, http.StatusOK, "User retrieved successfully", user)
}

func (h *userHandler) GetByPhone(w http.ResponseWriter, r *http.Request) {
	phone := chi.URLParam(r, "phone")

	user, err := h.userService.GetByPhone(phone)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve user", nil)
		return
	}

	response.Success(w, http.StatusOK, "User retrieved successfully", user)
}
