package router

import (
	"github.com/Morozov-rondo/restapi-template/internal/handlers"
	"github.com/go-chi/chi/v5"
)

// Функция регистрации маршрутов.
func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/1", h.Handler1.Command1)
	r.Get("/2", h.Handler1.Command2)
	r.Get("/3", h.Handler2.Command1)
	r.Get("/4", h.Handler2.Command2)

	return r
}
