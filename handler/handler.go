package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lsolniczek/api-mock/models"
)

type ServeHTTPImp func(w http.ResponseWriter, r *http.Request)

type serveHandler struct {
	ServeImp ServeHTTPImp
}

func (h *serveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ServeImp(w, r)
}

func NewServeHandler(stub models.APIStub) (http.Handler, error) {
	var body map[string]interface{}
	if err := json.Unmarshal([]byte(stub.Response.BodyJSON), &body); err != nil {
		return nil, err
	}
	serveImp := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range stub.Response.Headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(body)
	}
	return &serveHandler{serveImp}, nil
}
