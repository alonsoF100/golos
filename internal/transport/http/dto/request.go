package dto

// TODO добавить структуры данных для запросов

type UserRequest struct {
	Nickname string `json:"nickname" validate:"required,alphanum,min=3,max=12"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

type UserID struct {
	ID string `jsom:"id" validate:"required,uuid"`
}

type UserUpdate struct {
	ID       string `jsom:"id" validate:"required,uuid"`
	Nickname string `json:"nickname" validate:"required,alphanum,min=3,max=12"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}
type UserPatch struct {
	ID       string  `jsom:"id" validate:"required,uuid"`
	Nickname *string `json:"nickname,omitempty" validate:"alphanum,min=3,max=12"`
	Password *string `json:"password,omitempty" validate:"min=5,max=20"`
}
