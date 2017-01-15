package workmode

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/lsolniczek/mimus/models"
)

type Builder AppConfig

func (b *Builder) Start() error {
	return b.createNewProject()
}

func (b *Builder) createNewProject() error {
	newProjectPath := filepath.Join(b.ProjectsFilePath, b.ProjectName)
	_, err := os.Stat(newProjectPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(newProjectPath, 0777)
		} else {
			return err
		}
	}
	templatePath := filepath.Join(newProjectPath, "api-stub-template.json")
	content, err := projectCaseTemplate()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(templatePath, content, 0777)
}

func projectCaseTemplate() ([]byte, error) {
	return json.Marshal(models.APIStubTemplate())
}
