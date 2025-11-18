package service

type Repository interface {
	// TODO добавить контракты для интерфейса
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