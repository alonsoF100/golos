package service

import "github.com/alonsoF100/golos/internal/models"

type Repository interface {
	// TODO добавить контракты для интерфейса
	 UserExist(nicname string) (bool, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// TODO высрать методы сервисного слоя согласно контракту

func (s Service) CreateUser(nickname, password string) (*models.User, error) {
	// TODO  высрать функциональность
	// сходить в базу посмотреть есть ли пользователь
	// UserExist(nicname string) (bool, error)
	return nil, nil
}
