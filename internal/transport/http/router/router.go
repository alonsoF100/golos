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

	r.Route("/golos/elections", func(r chi.Router) {
		r.Post("/", rt.handlers.CreateElection)
		r.Get("/", rt.handlers.GetElections)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", rt.handlers.GetElection)
			r.Patch("/", rt.handlers.PatchElection)
			r.Delete("/", rt.handlers.DeleteElection)
		})
	})

	r.Route("/golos/vote_variants", func(r chi.Router) {
		r.Post("/", rt.handlers.CreateVoteVariant)
		r.Get("/", rt.handlers.GetVoteVariants)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", rt.handlers.GetVoteVariant)
			r.Put("/", rt.handlers.UpdateVoteVariant)
			r.Delete("/", rt.handlers.DeleteVoteVariant)
		})
	})

	r.Route("/golos/votes", func(r chi.Router) {
		r.Post("/", rt.handlers.CreateVote)
		r.Get("/", rt.handlers.GetUserVotes)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", rt.handlers.GetVote)
			r.Put("/", rt.handlers.DeleteVote)
			r.Delete("/", rt.handlers.PatchVote)
		})
	})

	return r
}
