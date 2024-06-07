package handlers

import (
	"github.com/Morozov-rondo/restapi-template/internal/handlers/handler1"
	"github.com/Morozov-rondo/restapi-template/internal/handlers/handler2"
	"github.com/Morozov-rondo/restapi-template/internal/service"
	"net/http"
)

type Handler struct {
	Handler1
	Handler2
}

func NewHandler(service *service.Service) *Handler {

	return &Handler{
		Handler1: handler1.NewHandler1(service.Service1),
		Handler2: handler2.NewHandler2(service.Service2),
	}
}

type Handler1 interface {
	Command1(w http.ResponseWriter, r *http.Request)
	Command2(w http.ResponseWriter, r *http.Request)
}

type Handler2 interface {
	Command1(w http.ResponseWriter, r *http.Request)
	Command2(w http.ResponseWriter, r *http.Request)
}
