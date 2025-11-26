package dto

// TODO добавить структуры данных для запросов

type UserRequest struct {
	Nickname string `json:"nickname" validate:"required,alphanum,min=3,max=12"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}