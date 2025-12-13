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
pattern: /golos/elections
method:  POST
info:    JSON in request body

succeed:
  - status code:   201 created
  - response body: JSON represented created election

failed:
  - status code:   400, 409, 500
  - response body: JSON with error + time
*/
func (h *Handler) CreateElection(w http.ResponseWriter, r *http.Request) {
	var req dto.ElectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	election, err := h.service.CreateElection(req.UserID, req.Name, req.Description)
	if err != nil {
		switch err {
		case apperrors.ErrElectionAlreadyExist:
			WriteJSON(w, http.StatusConflict, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusCreated, dto.NewElectionResponse(election))
}

/*
pattern: /golos/elections?limit=20&offset=0&nickname=alonso
method:  GET
info:    query (limit, offset, nickname)

succeed:
  - status code:   200 ok
  - response body: JSON represented elections

failed:
  - status code:   500
  - response body: JSON with error + time
*/
func (h *Handler) GetElections(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")
	nickName := query.Get("nickname")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	elections, err := h.service.GetElections(limit, offset, nickName)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	WriteJSON(w, http.StatusOK, dto.NewElectionsResponse(elections))
}

/*
pattern: /golos/elections/{id}
method:  GET
info:    UUID from pattern

succeed:
  - status code:   200 ok
  - response body: JSON represented election

failed:
  - status code:   400, 404, 500
  - response body: JSON with error + time
*/
func (h *Handler) GetElection(w http.ResponseWriter, r *http.Request) {
	var req dto.ElectionID
	req.ID = chi.URLParam(r, "id")

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	election, err := h.service.GetElection(req.ID)
	if err != nil {
		switch err {
		case apperrors.ErrElectionNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, dto.NewElectionResponse(election))
}

/*
pattern: /golos/elections/{id}
method:  DELETE
info:    UUID from pattern

succeed:
  - status code:   204 no content
  - response body: -

failed:
  - status code:   400, 404, 500
  - response body: JSON with error + time
*/
func (h *Handler) DeleteElection(w http.ResponseWriter, r *http.Request) {
	var req dto.ElectionID
	req.ID = chi.URLParam(r, "id")

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	err := h.service.DeleteElection(req.ID)
	if err != nil {
		switch err {
		case apperrors.ErrElectionNotFound:
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
pattern: /golos/elections/{id}
method:  PATCH
info:    UUID from pattern + JSON in request body

succeed:
  - status code:   200 ok
  - response body: JSON represented updated election

failed:
  - status code:   400, 404, 500
  - response body: JSON with error + time
*/
func (h *Handler) PatchElection(w http.ResponseWriter, r *http.Request) {
	var req dto.ElectionPatch
	req.ID = chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	election, err := h.service.PatchElection(req.ID, req.UserID, req.Name, req.Description)
	if err != nil {
		switch err {
		case apperrors.ErrElectionNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		case apperrors.ErrNothingToChange:
			WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, dto.NewElectionResponse(election))
}
