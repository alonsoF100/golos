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

// elections dto
type ElectionRequest struct {
	UserID      string `json:"user_id" validate:"required,uuid"`
	Name        string `json:"name" validate:"alphanum,min=3,max=50"`
	Description string `json:"description" validate:"required,min=3,max=100"`
}

type ElectionID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type ElectionPatch struct {
	ID          string  `json:"id" validate:"required,uuid"`
	UserID      *string `json:"user_id,omitempty" validate:"omitempty,uuid"`
	Name        *string `json:"name,omitempty" validate:"omitempty,alphanum,min=3,max=50"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=3,max=100"`
}

// Vote Variant DTOs
type VoteVariantRequest struct {
	ElectionID string `json:"election_id" validate:"required,uuid"`
	Name       string `json:"name" validate:"required,alphanum,min=1,max=50"`
}

type VoteVariantID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type VoteVariantUpdate struct {
	ID   string `json:"id" validate:"required,uuid"`
	Name string `json:"name" validate:"required,alphanum,min=1,max=50"`
}

type GetVoteVariantsRequest struct {
	ElectionID string `validate:"required,uuid"`
}

// vote dtos
type VoteRequest struct {
	VariantID string `json:"variat_id" validate:"required,uuid"`
	UserID    string `json:"user_id" validate:"required,uuid"`
}

type VoteID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type VotePatch struct {
	ID        string `json:"id"`
	VariantID string `json:"variant_id"`
	UserID    string `json:"user_id"`
}
