package workmode

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/lsolniczek/mimus/fileutil"
	"github.com/lsolniczek/mimus/models"
)

type ServerConfig struct {
	ProjectDetails ProjectDetails
}

func (sc *ServerConfig) BuildAPICases() ([]models.APICase, error) {
	if !sc.exist() {
		log.Fatalln("Project with name: " + sc.ProjectDetails.Name + " doesn't exist. Create it.")
	}

	// find project cases
	fileInfs, err := ioutil.ReadDir(sc.projectPath())
	if err != nil {
		return nil, err
	}

	// create array of ApiCase
	var apiCases []models.APICase
	for _, fileInfo := range fileInfs {
		if fileutil.ValidateFileExt(fileInfo.Name(), ".json") {
			casePath := filepath.Join(sc.projectPath(), fileInfo.Name())
			b, err := ioutil.ReadFile(casePath)
			if err != nil {
				log.Println(err)
				continue
			}
			apiCase := new(models.APICase)
			if err := json.Unmarshal(b, &apiCase); err != nil {
				log.Println(err)
				continue
			}
			apiCases = append(apiCases, *apiCase)
		}
	}
	return apiCases, nil
}

func (sc *ServerConfig) exist() bool {
	return fileutil.ProjectExists(sc.ProjectDetails.FilePath, sc.ProjectDetails.Name)
}

func (sc *ServerConfig) projectPath() string {
	return filepath.Join(sc.ProjectDetails.FilePath, sc.ProjectDetails.Name)
}
