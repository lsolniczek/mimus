package workmode

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lsolniczek/mimus/handler"
	"github.com/lsolniczek/mimus/models"
)

type server struct {
	APICases []models.APICase
}

func NewServer(apiCases []models.APICase) WorkMode {
	return &server{apiCases}
}

func (s *server) Start() error {
	router := s.buildRouter(s.APICases)
	return http.ListenAndServe(":8080", router)
}

func (s *server) buildRouter(apiCases []models.APICase) *mux.Router {
	r := mux.NewRouter()
	for _, apiCase := range apiCases {
		h := handler.NewServeHandler(apiCase)
		r.Handle(apiCase.Request.URL, h).Methods(apiCase.Request.Method)
	}
	return r
}
