package apimockmain

import (
	"flag"
	"log"
	"os"
	"os/user"
	"path/filepath"

	i "github.com/lsolniczek/api-mock/instance"
)

var fs = flag.NewFlagSet("api mock set", flag.ExitOnError)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fs.StringVar(&projectsFilePath, "path", apiMockConfigDir(), "Set custom path to project config directory")
}

func Run() i.Instance {
	cmd := os.Args[1]     // get command (run or new)
	name := os.Args[2]    // get command argument (project name)
	fs.Parse(os.Args[3:]) // parse arguments from 3th place to check flags

	// vaildate command and command argument
	if name == "" {
		log.Fatalln("You have to provide a projects name")
	}

	if (cmd != "new") && (cmd != "run") {
		log.Fatalln("There is no command " + cmd + ". You have to choose 'new' or 'run'")
	}

	// create create config directory if doesn't exist
	if err := createConfDirectory(projectsFilePath); err != nil {
		log.Fatalln(err)
	}

	// set AppConfig
	config := i.AppConfig{
		ProjectName:      name,
		ProjectsFilePath: projectsFilePath,
	}
	if cmd == "new" {
		return i.NewBuilder(config)
	}
	if cmd == "run" {
		return i.NewServer(config)
	}
	return nil
}

func createConfDirectory(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(path, 0777)
		} else {
			return err
		}
	}
	return nil
}

func apiMockConfigDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	return usr.HomeDir + string(filepath.Separator) + "apimock-config"
}

// application configurable parameters
var projectsFilePath string
