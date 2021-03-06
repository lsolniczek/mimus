package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lsolniczek/mimus/models"
)

type ServeHTTPImp func(w http.ResponseWriter, r *http.Request)

type serveHandler struct {
	ServeImp ServeHTTPImp
}

func (h *serveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ServeImp(w, r)
}

func NewServeHandler(stub models.APICase) http.Handler {
	serveImp := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range stub.Response.Headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stub.Response.BodyJSON)
	}
	return &serveHandler{serveImp}
}
