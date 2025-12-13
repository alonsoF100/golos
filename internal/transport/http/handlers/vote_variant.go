package handlers

import (
	"encoding/json"
	"net/http"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/transport/http/dto"
	"github.com/go-chi/chi/v5"
)

/*
pattern: /golos/vote_variants
method:  POST
info:    JSON in request body

succeed:
  - status code:   201 created
  - response body: JSON represented created vote variant

failed:
  - status code:   400, 409, 500
  - response body: JSON with error + time
*/
func (h *Handler) CreateVoteVariant(w http.ResponseWriter, r *http.Request) {
	var req dto.VoteVariantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	voteVariant, err := h.service.CreateVoteVariant(req.ElectionID, req.Name)
	if err != nil {
		switch err {
		case apperrors.ErrVoteVariantAlreadyExist:
			WriteJSON(w, http.StatusConflict, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusCreated, dto.NewVoteVariantResponse(voteVariant))
}

/*
pattern: /golos/vote-variants?election_id=elelnslksvmnspvmopsevmpoesvm
method:  GET
info:    electionID from query

succeed:
  - status code:   200 ok
  - response body: JSON represented vote variants

failed:
  - status code:   500
  - response body: JSON with error + time
*/
func (h *Handler) GetVoteVariants(w http.ResponseWriter, r *http.Request) {
	var req dto.GetVoteVariantsRequest
	query := r.URL.Query()
	req.ElectionID = query.Get("election_id")

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	voteVariants, err := h.service.GetVoteVariants(req.ElectionID)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	WriteJSON(w, http.StatusOK, dto.NewVoteVariantsResponse(voteVariants))
}

/*
pattern: /golos/vote_variants/{id}
method:  GET
info:    UUID from pattern

succeed:
  - status code:   200 ok
  - response body: JSON represented vote variant

failed:
  - status code:   400, 404, 500
  - response body: JSON with error + time
*/
func (h *Handler) GetVoteVariant(w http.ResponseWriter, r *http.Request) {
	var req dto.VoteVariantID
	req.ID = chi.URLParam(r, "id")

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	voteVariant, err := h.service.GetVoteVariant(req.ID)
	if err != nil {
		switch err {
		case apperrors.ErrVoteVariantNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, dto.NewVoteVariantResponse(voteVariant))
}

/*
pattern: /golos/vote_variants/{id}
method:  DELETE
info:    UUID from pattern

succeed:
  - status code:   204 no content
  - response body: -

failed:
  - status code:   400, 404, 500
  - response body: JSON with error + time
*/
func (h *Handler) DeleteVoteVariant(w http.ResponseWriter, r *http.Request) {
	var req dto.VoteVariantID
	req.ID = chi.URLParam(r, "id")

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	err := h.service.DeleteVoteVariant(req.ID)
	if err != nil {
		switch err {
		case apperrors.ErrVoteVariantNotFound:
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
pattern: /golos/vote_variants/{id}
method:  PUT
info:    UUID from pattern + JSON in request body

succeed:
  - status code:   200 ok
  - response body: JSON represented updated vote variant

failed:
  - status code:   400, 404, 500
  - response body: JSON with error + time
*/
func (h *Handler) UpdateVoteVariant(w http.ResponseWriter, r *http.Request) {
	var req dto.VoteVariantUpdate
	req.ID = chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	voteVariant, err := h.service.UpdateVoteVariant(req.ID, req.Name)
	if err != nil {
		switch err {
		case apperrors.ErrVoteVariantNotFound:
			WriteJSON(w, http.StatusNotFound, dto.NewErrorResponse(err))
			return
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}

	WriteJSON(w, http.StatusOK, dto.NewVoteVariantResponse(voteVariant))
}
