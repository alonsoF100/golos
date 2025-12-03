package handlers

import (
	"encoding/json"
	"net/http"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/transport/http/dto"
	"github.com/go-chi/chi/v5"
)

/*
pattern: /golos/votes
method:  POST
info:    JSON in request body

succeed:

	-status code:   201 created
	-response body: JSON represented created vote

failed:

	-status code:   400, 409, 500
	-response body: JSON with error + time
*/
func (h *Handler) CreateVote(w http.ResponseWriter, r *http.Request) {
	var req dto.VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	vote, err := h.service.CreateVote(req.UserID, req.VariantID)
	if err != nil {
		switch err {
		case apperrors.ErrVoteAlreadyExist:
			WriteJSON(w, http.StatusConflict, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusCreated, dto.NewVoteResponse(vote))
}

/*
pattern: /golos/votes/{id}
method:  GET
info:    UUID from pattern

succeed:

	-status code:   200 ok
	-response body: JSON represented vote

failed:

	-status code:   400, 404, 500
	-response body: JSON with error + time
*/
func (h *Handler) GetVote(w http.ResponseWriter, r *http.Request) {
	var req dto.VoteID

	req.ID = chi.URLParam(r, "id")

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	vote, err := h.service.GetVote(req.ID)
	if err != nil {
		switch err {
		case apperrors.ErrVoteNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, dto.NewVoteResponse(vote))
}

/*
pattern: /golos/votes/{id}
method:  DELETE
info:    UUID from pattern

succeed:

	-status code:   204 no content
	-response body: -

failed:

	-status code:   400, 404, 500
	-response body: JSON with error + time
*/
func (h *Handler) DeleteVote(w http.ResponseWriter, r *http.Request) {
	var req dto.VoteID

	req.ID = chi.URLParam(r, "id")

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	err := h.service.DeleteVote(req.ID)
	if err != nil {
		switch err {
		case apperrors.ErrVoteNotFound:
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
pattern: /golos/votes/{id}
method:  PATCH
info:    UUID from pattern + JSON in request body

succeed:

	-status code:   200 ok
	-response body: JSON represented updated vote

failed:

	-status code:   400, 404, 500
	-response body: JSON with error + time
*/
func (h *Handler) PatchVote(w http.ResponseWriter, r *http.Request) {
	var req dto.VotePatch

	req.ID = chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	vote, err := h.service.PatchVote(req.ID, &req.UserID, &req.VariantID)
	if err != nil {
		switch err {
		case apperrors.ErrVoteNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		case apperrors.ErrUserNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		case apperrors.ErrElectionNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, dto.NewVoteResponse(vote))
}
