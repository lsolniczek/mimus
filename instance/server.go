package instance

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/lsolniczek/api-mock/fileUtil"
	"github.com/lsolniczek/api-mock/handler"
	"github.com/lsolniczek/api-mock/models"
	"path"
)

type server struct {
	Config AppConfig
}

func NewServer(config AppConfig) Instance {
	return &server{config}
}

func (s *server) Start() {
	log.Fatalln(s.runServer())
}

func (s *server) runServer() error {
	if !s.exist() {
		return errors.New("Project with name: " + s.Config.ProjectName + " doesn't exist. Create it.")
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

func (s *server) exist() bool {
	return fileUtil.ProjectExist(s.Config.ProjectsFilePath, s.Config.ProjectName)
}

func (s *server) templatesNames() ([]string, error) {
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

func (s *server) buildServer(requestNames []string) error {
	r := mux.NewRouter()
	for _, reqName := range requestNames {
		path := s.projectPath() + string(filepath.Separator) + reqName
		apiStub, err := apiStubFromTemplateFile(path)
		if err != nil {
			return err
		}
		h, err := handler.NewHandler(apiStub)
		if err != nil {
			return err
		}
		r.HandleFunc(h.URLPath, h.HandlerFunc).Methods(h.HTTPMethod)
	}
	return http.ListenAndServe(":8080", r)
}

func (s *server) projectPath() string {
	return s.Config.ProjectsFilePath + string(filepath.Separator) + s.Config.ProjectName
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
