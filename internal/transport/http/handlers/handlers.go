package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alonsoF100/golos/internal/models"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	// User service methods
	CreateUser(nickname, password string) (*models.User, error)
	GetUsers(limit, offset int) ([]*models.User, error)
	GetUser(uuid string) (*models.User, error)
	UpdateUser(uuid, nickname, password string) (*models.User, error)
	DeleteUser(uuid string) error
	PatchUser(uuid string, nickname, password *string) (*models.User, error)

	// Election service methods
	CreateElection(userID string, name string, description string) (*models.Election, error)
	GetElections(limit, offset int, nickname string) ([]*models.Election, error)
	GetElection(uuid string) (*models.Election, error)
	DeleteElection(uuid string) error
	PatchElection(uuid string, userID, name, description *string) (*models.Election, error)

	// Vote-variant service methods
	CreateVoteVariant(electionID, name string) (*models.VoteVariant, error)
	GetVoteVariants(electionID string) ([]*models.VoteVariant, error)
	GetVoteVariant(uuid string) (*models.VoteVariant, error)
	DeleteVoteVariant(uuid string) error
	UpdateVoteVariant(uuid string, name string) (*models.VoteVariant, error)

	// Vote service methods
	CreateVote(userID, voteVariantID string) (*models.Vote, error)
	GetVote(voteID string) (*models.Vote, error)
	GetUserVotes(userID string) (*[]models.Vote, error)
	GetVariantVotes(voteVariantID string) (*[]models.Vote, error)
	DeleteVote(voteID string) error
	PatchVote(voteID string, userID, voteVariantID *string) (*models.Vote, error)
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
