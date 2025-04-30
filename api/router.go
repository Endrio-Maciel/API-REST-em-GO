package api

import (
	"api_rest/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Service *service.Application
}

func (app *Application) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", app.handleGetAllUsers)
		r.Post("/", app.handleCreateUser)
		r.Get("/{id}", app.handleGetUserById)
		r.Put("/{id}", app.handleUpdateUser)
		r.Delete("/{id}", app.handleDeleteUser)
	})

	return r
}
