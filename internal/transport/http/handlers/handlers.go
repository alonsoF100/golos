package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	apperrors "github.com/alonsoF100/golos/internal/erorrs"
	"github.com/alonsoF100/golos/internal/models"
	"github.com/alonsoF100/golos/internal/transport/http/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	// User methods
	CreateUser(nickname, password string) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUser(uuid string) (*models.User, error)
	UpdateUser(uuid, nickname, password string) (*models.User, error)
	DeleteUser(uuid string) error
	PatchUser(uuid string, nickname, password *string) (*models.User, error)

	// Election methods
	CreateElection(userID string, name string, description *string) (*models.Election, error)
	GetElections() ([]*models.Election, error)
	GetElection(uuid string) (*models.Election, error)
	DeleteElection(uuid string) error
	PatchElection(uuid string, userID, name, description *string) (*models.Election, error)

	// Vote Variant methods
	CreateVoteVariant(electionID, name string) (*models.VoteVariant, error)
	GetVoteVariants() ([]*models.VoteVariant, error)
	GetVoteVariant(uuid string) (*models.VoteVariant, error)
	DeleteVoteVariant(uuid string) error
	UpdateVoteVariant(uuid string, name string) (*models.VoteVariant, error)
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

	if err := h.Validator.Struct(req); err != nil {
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
pattern: /golos/users
method:  GET
info:    -

succeed:

	-status code:   200 ok
	-response body: JSON represented users

failed:

	-status code:   500
	-response body: JSON with error + time
*/
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// 1. Cходить в сервисный слой за пользователями(вернется слайс + ошибка)
	users, err := h.service.GetUsers()
	// 2. Обработать ошибки
	if err != nil {
		switch err {
		default:
			WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}
	// 3. Собрать(сделать функцию helper в dto response, которая переберт полученный слайс) и
	// отправить ответ, содержащий всех пользователей + статус код
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
	// 1. Создать dto где будет лежать айдишник пользователя
	var req dto.UserID
	// 2. Записать в dto айдишник
	req.ID = chi.URLParam(r, "id")
	// 3. Запарсить dto через validator/v10
	if err := h.Validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}
	// 4. Сходить в сервисный слой за пользователем(вернется пользователь + ошибка)
	user, err := h.service.GetUser(req.ID)
	// 5. обработать ошибки(пользователь не найден + ошибка сервака)
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
	// 6. Собрать ответ и отправить его вместе со статус кодом
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
	// 1. Создать dto где будет лежать айдишник пользователя + ВООБЩЕ ВСЕ КРОМЕ СЛУЖЕБНЫХ поля для записи
	var req dto.UserUpdate
	// 2. Записать в dto айдишник
	req.ID = chi.URLParam(r, "id")
	// 3. Записать в dto полученные в json поля
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}
	// 4. Запарсить dto через validator/v10
	if err := h.Validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}
	// 5. Сходить в сервисный слой для обновления пользователя(вернется пользователь + ошибка)
	user, err := h.service.UpdateUser(req.ID, req.Nickname, req.Password)
	// 6. обработать ошибки(пользователь не найден + ошибка сервака)
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
	// 7. Собрать ответ и отправить его вместе со статус кодом
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
	// 1. Создать dto где будет лежать айдишник пользователя
	var req dto.UserID
	// 2. Записать в dto айдишник
	req.ID = chi.URLParam(r, "id")
	// 3. Запарсить dto через validator/v10
	if err := h.Validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}
	// 4. Сходить в сервисный слой(удалить пользователя) вернется(ошибка)
	err := h.service.DeleteUser(req.ID)
	// 5. обработать ошибки(конфликт + ошибка сервака)
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
	// 6. В случае успех выернуть 204 no content
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
	// 1. Создать dto где будет лежать айдишник пользователя + ТОЛЬКО УКАЗАННЫЕ ПОЛЯ( ЧАСТИЧНОЕ ОБНОВЛЕНИЕ PATCH)
	var req dto.UserPatch
	// для записи ТУТ ЕЩЕ НАДО ПОМНИТЬ ЧТО ПОЛЯ МОГУ БЫТЬ НЕОБЯЗАТЕЛЬНЫМИ ПАРОЛЬ И НИКНЕЙМ ТАК ЧТО ОНИ БУДЕТ ТИПА *string
	// Чтобы если ничо не пришло выдавать nil в слой сервиса, чтобы там по кайфу обработать чо как газ газ
	// 2. Записать в dto айдишник
	req.ID = chi.URLParam(r, "id")
	// 3. Записать в dto полученные в json поля
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}
	// 4. Запарсить dto через validator/v10
	if err := h.Validator.Struct(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}
	// 5. Сходить в сервисный слой для обновления пользователя(вернется пользователь + ошибка)
	user, err := h.service.PatchUser(req.ID, req.Nickname, req.Password)
	// 6. обработать ошибки(пользователь не найден + ошибка сервака)
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
	// 7. Собрать ответ и отправить его вместе со статус кодом
	WriteJSON(w, http.StatusOK, dto.NewUserResponse(user))
}

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

	if err := h.Validator.Struct(req); err != nil {
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
pattern: /golos/elections
method:  GET
info:    -

succeed:
  - status code:   200 ok
  - response body: JSON represented elections

failed:
  - status code:   500
  - response body: JSON with error + time
*/
func (h *Handler) GetElections(w http.ResponseWriter, r *http.Request) {
	elections, err := h.service.GetElections()
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

	if err := h.Validator.Struct(req); err != nil {
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

	if err := h.Validator.Struct(req); err != nil {
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

	if err := h.Validator.Struct(req); err != nil {
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

// Vote Variant Handlers //

/*
pattern: /golos/vote-variants
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

	if err := h.Validator.Struct(req); err != nil {
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
pattern: /golos/vote-variants
method:  GET
info:    -

succeed:
  - status code:   200 ok
  - response body: JSON represented vote variants

failed:
  - status code:   500
  - response body: JSON with error + time
*/
func (h *Handler) GetVoteVariants(w http.ResponseWriter, r *http.Request) {
	voteVariants, err := h.service.GetVoteVariants()
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	WriteJSON(w, http.StatusOK, dto.NewVoteVariantsResponse(voteVariants))
}

/*
pattern: /golos/vote-variants/{id}
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

	if err := h.Validator.Struct(req); err != nil {
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
pattern: /golos/vote-variants/{id}
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

	if err := h.Validator.Struct(req); err != nil {
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
pattern: /golos/vote-variants/{id}
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

	if err := h.Validator.Struct(req); err != nil {
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
