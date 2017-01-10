package instance

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/lsolniczek/api-mock/models"
)

type builder struct {
	Config AppConfig
}

func NewBuilder(config AppConfig) Instance {
	return &builder{config}
}

func (b *builder) Start() {
	log.Fatalln(b.createNewProject())
}

func (b *builder) createNewProject() error {
	newProjectPath := b.Config.ProjectsFilePath + string(filepath.Separator) + b.Config.ProjectName
	_, err := os.Stat(newProjectPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(newProjectPath, 0777)
		} else {
			return err
		}
	}
	templatePath := newProjectPath + string(filepath.Separator) + "api-stub-template.json"
	content, err := projectCaseTemplate()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(templatePath, content, 0777)
}

func projectCaseTemplate() ([]byte, error) {
	apiStub := models.APIStubTemplate()
	return json.Marshal(apiStub)
}
