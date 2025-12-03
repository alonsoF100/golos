package dto

import (
	"time"

	"github.com/alonsoF100/golos/internal/models"
)

type ErrorResponse struct {
	Error     string    `json:"error"`
	TimeStamp time.Time `json:"timestamp"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error:     err.Error(),
		TimeStamp: time.Now(),
	}
}

type UserResponse struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Password:  user.Password,
		CreatedAt: user.UpdatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UsersResponse struct {
	Users []*UserResponse
}

func NewUsersResponse(users []*models.User) UsersResponse {
	responseUsers := UsersResponse{
		Users: make([]*UserResponse, 0, len(users)),
	}

	for _, user := range users {
		temp := &UserResponse{
			ID:        user.ID,
			Nickname:  user.Nickname,
			Password:  user.Password,
			CreatedAt: user.UpdatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		responseUsers.Users = append(responseUsers.Users, temp)
	}

	return responseUsers
}

// election dto
type ElectionResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewElectionResponse(election *models.Election) ElectionResponse {
	return ElectionResponse{
		ID:          election.ID,
		UserID:      election.UserID,
		Name:        election.Name,
		Description: election.Description,
		CreatedAt:   election.CreatedAt,
		UpdatedAt:   election.UpdatedAt,
	}
}

type ElectionsResponse struct {
	Elections []*ElectionResponse
}

func NewElectionsResponse(elections []*models.Election) ElectionsResponse {
	responseElections := ElectionsResponse{
		Elections: make([]*ElectionResponse, 0, len(elections)),
	}
	for _, election := range elections {
		temp := &ElectionResponse{
			ID:          election.ID,
			UserID:      election.UserID,
			Name:        election.Name,
			Description: election.Description,
			CreatedAt:   election.CreatedAt,
			UpdatedAt:   election.UpdatedAt,
		}
		responseElections.Elections = append(responseElections.Elections, temp)
	}
	return responseElections
}

// Vote Variant Responses
type VoteVariantResponse struct {
	ID         string    `json:"id"`
	ElectionID string    `json:"election_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewVoteVariantResponse(voteVariant *models.VoteVariant) VoteVariantResponse {
	return VoteVariantResponse{
		ID:         voteVariant.ID,
		ElectionID: voteVariant.ElectionID,
		Name:       voteVariant.Name,
		CreatedAt:  voteVariant.CreatedAt,
		UpdatedAt:  voteVariant.UpdatedAt,
	}
}

type VoteVariantsResponse struct {
	VoteVariants []*VoteVariantResponse
}

func NewVoteVariantsResponse(voteVariants []*models.VoteVariant) VoteVariantsResponse {
	responseVariants := VoteVariantsResponse{
		VoteVariants: make([]*VoteVariantResponse, 0, len(voteVariants)),
	}
	for _, variant := range voteVariants {
		temp := &VoteVariantResponse{
			ID:         variant.ID,
			ElectionID: variant.ElectionID,
			Name:       variant.Name,
			CreatedAt:  variant.CreatedAt,
			UpdatedAt:  variant.UpdatedAt,
		}
		responseVariants.VoteVariants = append(responseVariants.VoteVariants, temp)
	}
	return responseVariants
}

type VoteResponse struct {
	ID        string    `json:"id"`
	VariantID string    `json:"variant_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVoteResponse(vote *models.Vote) VoteResponse {
	return VoteResponse{
		ID:        vote.ID,
		VariantID: vote.VariantID,
		UserID:    vote.UserID,
		CreatedAt: vote.CreatedAt,
		UpdatedAt: vote.UpdatedAt,
	}
}
