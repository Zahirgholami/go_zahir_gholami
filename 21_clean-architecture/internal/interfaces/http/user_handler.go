package http

import (
	"encoding/json"
	"net/http"
	"project/internal/domain/model"
	"project/internal/usecase"
)

type UserHandler struct {
	UseCase *usecase.UserUseCase
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	err := h.UseCase.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
