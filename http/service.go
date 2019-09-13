package http

import (
	_ "net/http"

	"github.com/gorilla/mux"
)

type WebService struct {
	r *mux.Router
}

func NewWebService(address string) *WebService {
	return &WebService{
		r: mux.NewRouter(),
	}
}
