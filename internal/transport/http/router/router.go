package router

import (
	"github.com/alonsoF100/golos/internal/transport/http/handlers"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	handlers *handlers.Handler
}

func New(handlers *handlers.Handler) *Router {
	return &Router{
		handlers: handlers,
	}
}

func (rt Router) Setup() *chi.Mux {
	r := chi.NewRouter()

	// TODO настроитьмрашруты
	// TODO настроить конфигурацию middleware

	// r.Route("/golos/")

	return r
}