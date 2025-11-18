package handlers

type Service interface {
	// TODO добавить контракты для интерфейса

}

type Handler struct {
	service Service
}

func New(service Service) *Handler{
	return &Handler{
		service: service,
	}
}

// TODO добавить хендлеры + описапние контратов для документации
/*
контракт
func (handler * Handler) название (w http.ResponseWriter, r *http.Request) {
}
*/

