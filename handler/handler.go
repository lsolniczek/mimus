package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lsolniczek/api-mock/models"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type handler struct {
	HandlerFunc HandlerFunc
	HTTPMethod  string
	URLPath     string
}

func NewHandler(stub models.APIStub) (*handler, error) {
	var body map[string]interface{}
	if err := json.Unmarshal([]byte(stub.Response.BodyJSON), &body); err != nil {
		return nil, err
	}
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(body)
	}
	return &handler{
		HandlerFunc: handlerFunc,
		HTTPMethod:  stub.Request.Method,
		URLPath:     stub.Request.URL,
	}, nil
}
