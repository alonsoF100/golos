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

	r.Route("/golos/users", func(r chi.Router) {
		r.Post("/", rt.handlers.CreateUser)
		r.Get("/", rt.handlers.GetUsers)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", rt.handlers.GetUser)
			r.Put("/", rt.handlers.UpdateUser)
			r.Patch("/", rt.handlers.PatchUser)
			r.Delete("/", rt.handlers.DeleteUser)
		})
	})

	return r
}
