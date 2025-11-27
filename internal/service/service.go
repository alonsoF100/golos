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

// TODO выдавить документацию к функциям малую 

func (s Service) CreateUser(nickname, password string) (*models.User, error) {
	// TODO  высрать функциональность
	// сходить в базу посмотреть есть ли пользователь
	// UserExist(nicname string) (bool, error)
	return nil, nil
}

func (s Service) GetUsers() ([]*models.User, error) {
	// TODO высрать функциональность

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil, nil
}

func (s Service) GetUser(uuid string) (*models.User, error) {
	// TODO высрать функциональность

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil, nil
}

func (s Service) UpdateUser(uuid, nickname, password string) (*models.User, error) {
	// TODO высрать функциональность

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil, nil
}

func (s Service) DeleteUser(uuid string) error {
	// TODO высрать функциональность

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil
}

func (s Service) PatchUser(uuid string, nickname, password *string) (*models.User, error) {
	// TODO высрать функциональность

	// Пока загшлука на вывод, чтобы красным не гадило
	return nil, nil
}
