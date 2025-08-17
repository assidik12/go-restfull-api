package handler

import (
	"encoding/json"
	"net/http"

	"github.com/assidik12/go-restfull-api/internal/delivery/http/dto"
	"github.com/assidik12/go-restfull-api/internal/pkg/response"
	"github.com/assidik12/go-restfull-api/internal/service"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	service service.UserService
}

// Implement methods for UserHandler as needed
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	user, err := h.service.Register(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.Created(w, user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	token, err := h.service.Login(r.Context(), req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.OK(w, token)
}
