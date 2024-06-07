package handler2

import (
	"fmt"
	"github.com/Morozov-rondo/restapi-template/internal/service"
	"net/http"
)

func NewHandler2(service service.Service2) *Handler2 {
	return &Handler2{
		service: service,
	}
}

type Handler2 struct {
	service service.Service2
}

func (h *Handler2) Command1(w http.ResponseWriter, r *http.Request) {

	// TODO: implement

	fmt.Fprint(w, "Handler1/Command1")
}

func (h *Handler2) Command2(w http.ResponseWriter, r *http.Request) {

	// TODO: implement

	fmt.Fprint(w, "Handler1/Command2")
}
