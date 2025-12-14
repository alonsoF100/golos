package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/transport/http/dto"
	"github.com/go-chi/chi/v5"
)

/*
pattern: /golos/users
method:  POST
info:    JSON in request body

succeed:

	-status code:   201 created
	-response body: JSON represented created user

failed:

	-status code:   400, 409, 500
	-response body: JSON with error + time
*/
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
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

	WriteJSON(w, http.StatusCreated, dto.NewUserResponse(user))
}

/*
pattern: /golos/users?limit=20&offset=20
method:  GET
info:    query (limit, offset)

succeed:

	-status code:   200 ok
	-response body: JSON represented users

failed:

	-status code:   500
	-response body: JSON with error + time
*/
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	users, err := h.service.GetUsers(limit, offset)
	if err != nil {
		switch err {
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, dto.NewUsersResponse(users))
}

/*
pattern: /golos/users/{id}
method:  GET
info:    UUID from pattern

succeed:

	-status code:   200 ok
	-response body: JSON represented user

failed:

	-status code:   400, 404, 500
	-response body: JSON with error + time
*/
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserID

	req.ID = chi.URLParam(r, "id")

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	user, err := h.service.GetUser(req.ID)
	if err != nil {
		switch err {
		case apperrors.ErrUserNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, dto.NewUserResponse(user))
}

/*
pattern: /golos/users/{id}
method:  PUT
info:    UUID from pattern + JSON in request body

succeed:

	-status code:   200 ok
	-response body: JSON represented updated user

failed:

	-status code:   400, 404, 500
	-response body: JSON with error + time
*/
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserUpdate

	req.ID = chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	user, err := h.service.UpdateUser(req.ID, req.Nickname, req.Password)
	if err != nil {
		switch err {
		case apperrors.ErrUserNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, user)
}

/*
pattern: /golos/users/{id}
method:  DELETE
info:    UUID from pattern

succeed:

	-status code:   204 no content
	-response body: -

failed:

	-status code:   400, 404, 500
	-response body: JSON with error + time
*/
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserID

	req.ID = chi.URLParam(r, "id")

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	err := h.service.DeleteUser(req.ID)
	if err != nil {
		switch err {
		case apperrors.ErrUserNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusNoContent, nil)
}

/*
pattern: /golos/users/{id}
method:  PATCH
info:    UUID from pattern + JSON in request body

succeed:

	-status code:   200 ok
	-response body: JSON represented updated user

failed:

	-status code:   400, 404, 500
	-response body: JSON with error + time
*/
func (h *Handler) PatchUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserPatch

	req.ID = chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	user, err := h.service.PatchUser(req.ID, req.Nickname, req.Password)
	if err != nil {
		switch err {
		case apperrors.ErrUserNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, dto.NewUserResponse(user))
}
