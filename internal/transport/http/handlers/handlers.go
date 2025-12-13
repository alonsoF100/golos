package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alonsoF100/golos/internal/models"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	CreateUser(nickname, password string) (*models.User, error)
	GetUsers(limit, offset int) ([]*models.User, error)
	GetUser(uuid string) (*models.User, error)
	UpdateUser(uuid, nickname, password string) (*models.User, error)
	DeleteUser(uuid string) error
	PatchUser(uuid string, nickname, password *string) (*models.User, error)
}

type ElectionService interface {
	CreateElection(userID string, name string, description string) (*models.Election, error)
	GetElection(uuid string) (*models.Election, error)
	DeleteElection(uuid string) error
	PatchElection(uuid string, userID, name, description *string) (*models.Election, error)
}

type ElectionQueryService interface {
	GetElections(limit, offset int, nickname string) ([]*models.Election, error)
}

type VoteVariantService interface {
	CreateVoteVariant(electionID, name string) (*models.VoteVariant, error)
	GetVoteVariants(electionID string) ([]*models.VoteVariant, error)
	GetVoteVariant(uuid string) (*models.VoteVariant, error)
	DeleteVoteVariant(uuid string) error
	UpdateVoteVariant(uuid string, name string) (*models.VoteVariant, error)
}

type VoteService interface {
	CreateVote(userID, voteVariantID string) (*models.Vote, error)
	GetVote(voteID string) (*models.Vote, error)
	GetUserVotes(userID string) (*[]models.Vote, error)
	GetVariantVotes(voteVariantID string) (*[]models.Vote, error)
	DeleteVote(voteID string) error
	PatchVote(voteID string, userID, voteVariantID *string) (*models.Vote, error)
}

type Service interface {
	UserService
	ElectionService
	ElectionQueryService
	VoteVariantService
	VoteService
}

type Handler struct {
	service   Service
	validator *validator.Validate
}

func New(service Service) *Handler {
	return &Handler{
		service:   service,
		validator: validator.New(),
	}
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("error: %v, time: %v\n", err.Error(), time.Now())
	}
}
