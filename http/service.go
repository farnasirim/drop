package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebService struct {
	r *mux.Router
}

func NewWebService(address string) {
	return WebService{
		r: mux.NewRouter(),
	}
}
