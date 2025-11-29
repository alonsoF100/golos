package dto

type UserRequest struct {
	Nickname string `json:"nickname" validate:"required,alphanum,min=3,max=12"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

type UserID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type UserUpdate struct {
	ID       string `json:"id" validate:"required,uuid"`
	Nickname string `json:"nickname" validate:"required,alphanum,min=3,max=12"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

type UserPatch struct {
	ID       string  `json:"id" validate:"required,uuid"`
	Nickname *string `json:"nickname,omitempty" validate:"omitempty,alphanum,min=3,max=12"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=5,max=20"`
}

type ElectionRequest struct {
	UserID      string  `json:"user_id" validate:"required,uuid"`
	Name        string  `json:"name" validate:"alphanum,min=3,max=50"`
	Description *string `json:"nickname,omitempty" validate:"omitempty,min=3,max=100"`
}
