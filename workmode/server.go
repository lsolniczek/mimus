package workmode

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/lsolniczek/api-mock/fileutil"
	"github.com/lsolniczek/api-mock/handler"
	"github.com/lsolniczek/api-mock/models"
)

type Server AppConfig

func (s *Server) Start() error {
	return s.runServer()
}

func (s *Server) runServer() error {
	if !s.exist() {
		return errors.New("Project with name: " + s.ProjectName + " doesn't exist. Create it.")
	}
	requestsNames, err := s.templatesNames()
	if err != nil {
		return err
	}
	if err := s.buildServer(requestsNames); err != nil {
		return err
	}
	return nil
}

func (s *Server) exist() bool {
	return fileutil.ProjectExists(s.ProjectsFilePath, s.ProjectName)
}

func (s *Server) templatesNames() ([]string, error) {
	var names []string
	projectPath := s.projectPath()
	fis, err := ioutil.ReadDir(projectPath)
	if err != nil {
		return nil, err
	}
	for _, fi := range fis {
		if !fi.IsDir() && path.Ext(fi.Name()) == ".json" {
			names = append(names, fi.Name())
		}
	}
	return names, nil
}

func (s *Server) buildServer(requestNames []string) error {
	r := mux.NewRouter()
	for _, reqName := range requestNames {
		path := filepath.Join(s.projectPath(), reqName)
		apiStub, err := apiStubFromTemplateFile(path)
		if err != nil {
			return err
		}
		h, err := handler.NewServeHandler(apiStub)
		if err != nil {
			return err
		}
		r.Handle(apiStub.Request.URL, h).Methods(apiStub.Request.Method)
	}
	return http.ListenAndServe(":8080", r)
}

func (s *Server) projectPath() string {
	return filepath.Join(s.ProjectsFilePath, s.ProjectName)
}

func apiStubFromTemplateFile(path string) (models.APIStub, error) {
	var stub models.APIStub
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return stub, err
	}
	if err := json.Unmarshal(b, &stub); err != nil {
		return stub, err
	}
	return stub, nil
}
