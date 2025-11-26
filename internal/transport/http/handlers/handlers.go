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
pattern: /golos/users
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
pattern: /users
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
	// TODO реализовать хендлер для получения всех пользователей

	// 1. Cходить в сервисный слой за пользователями(вернется слайс + ошибка)
	// 2. Обработать ошибки
	// 3. Собрать(сделать функцию helper в dto response, которая переберт полученный слайс) и
	// отправить ответ, содержащий всех пользователей + статус код
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
	// TODO реализовать хендлер для получения пользователя по айдишнику

	// 1. Создать dto где будет лежать айдишник пользователя
	// 2. Записать в dto айдишник
	// 3. Запарсить dto через validator/v10
	// 4. Сходить в сервисный слой за пользователем(вернется пользователь + ошибка)
	// 5. обработать ошибки(пользователь не найден + ошибка сервака)
	// 6. Собрать ответ и отправить его вместе со статус кодом
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
	// TODO реализовать хендлер по обновлению всех полей пользователя по айдишнику

	// 1. Создать dto где будет лежать айдишник пользователя + ВООБЩЕ ВСЕ КРОМЕ СЛУЖЕБНЫХ поля для записи
	// 2. Записать в dto айдишник
	// 3. Записать в dto полученные в json поля
	// 4. Запарсить dto через validator/v10
	// 5. Сходить в сервисный слой для обновления пользователя(вернется пользователь + ошибка)
	// 6. обработать ошибки(пользователь не найден + ошибка сервака)
	// 7. Собрать ответ и отправить его вместе со статус кодом

}

/*
pattern: /golos/users/{id}
method:  DELETE
info:    UUID from pattern

succeed:

	-status code:   204 no content
	-response body: -

failed:

	-status code:   400, 409, 500
	-response body: JSON with error + time
*/
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO реализовать удаление пользователя по айдишнику

	// 1. Создать dto где будет лежать айдишник пользователя
	// 2. Записать в dto айдишник
	// 3. Запарсить dto через validator/v10
	// 4. Сходить в сервисный слой(удалить пользователя) вернется(ошибка)
	// 5. обработать ошибки(конфликт + ошибка сервака)
	// 6. В случае успех выернуть 204 no content
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
func (h *Handler) PathUser(w http.ResponseWriter, r *http.Request) {
	// TODO реализовать частичное обновления пользователя по айдишнику + указанным в json полям

	// 1. Создать dto где будет лежать айдишник пользователя + ТОЛЬКО УКАЗАННЫЕ ПОЛЯ( ЧАСТИЧНОЕ ОБНОВЛЕНИЕ PATCH) для записи
	// 2. Записать в dto айдишник
	// 3. Записать в dto полученные в json поля
	// 4. Запарсить dto через validator/v10
	// 5. Сходить в сервисный слой для обновления пользователя(вернется пользователь + ошибка)
	// 6. обработать ошибки(пользователь не найден + ошибка сервака)
	// 7. Собрать ответ и отправить его вместе со статус кодом
}
