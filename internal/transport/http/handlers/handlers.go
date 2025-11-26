package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/alonsoF100/golos/internal/transport/http/dto"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	// TODO добавить контракты для интерфейса
	CreateUser(nickname, password string) (*models.User, error)
}

type Handler struct {
	service   Service
	Validator *validator.Validate
}

func New(service Service) *Handler {
	return &Handler{
		service:   service,
		Validator: validator.New(),
	}
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("error: %v, time: %v\n", err.Error(), time.Now())
	}
}

// TODO добавить хендлеры + описапние контратов для документации
/*
method:  POST
info:    JSON in request body

succeed:

	-status code:   201 created
	-response body: JSON represented created product

failed:

	-status code:   400, 409, 500
	-response body: JSON with error + time
*/
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
	}

	if err := h.Validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
	}

	user, err := h.service.CreateUser(req.Nickname, req.Password)
	if err != nil {
		switch err {
		case apperrors.ErrUserAlreadyExist:
			WriteJSON(w, http.StatusConflict, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusInternalServerError, dto.NewUserResponse(user))
}
