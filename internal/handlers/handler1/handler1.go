package handler1

import (
	"fmt"
	"github.com/Morozov-rondo/restapi-template/internal/service"
	"net/http"
)

func NewHandler1(service service.Service1) *Handler1 {
	return &Handler1{
		service: service,
	}
}

type Handler1 struct {
	service service.Service1
}

func (h *Handler1) Command1(w http.ResponseWriter, r *http.Request) {

	// TODO: implement

	fmt.Fprint(w, "Handler1/Command1")
}

func (h *Handler1) Command2(w http.ResponseWriter, r *http.Request) {

	// TODO: implement

	fmt.Fprint(w, "Handler1/Command2")
}
