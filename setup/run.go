package setup

import (
	"flag"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/lsolniczek/api-mock/workmode"
)

var fs = flag.NewFlagSet("api mock set", flag.ExitOnError)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fs.StringVar(&projectsFilePath, "path", apiMockConfigDir(), "Set custom path to project config directory")
}

func Run() workmode.WorkMode {
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
	createConfDirectory(projectsFilePath)

	// set AppConfig
	config := workmode.AppConfig{
		ProjectName:      name,
		ProjectsFilePath: projectsFilePath,
	}
	if cmd == "new" {
		b := workmode.Builder(config)
		return &b
	}
	if cmd == "run" {
		s := workmode.Server(config)
		return &s
	}
	return nil
}

func createConfDirectory(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(path, 0777)
		} else {
			log.Fatalln(err)
		}
	}
}

func apiMockConfigDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	return filepath.Join(usr.HomeDir, "apimock-config")
}

// application configurable parameters
var projectsFilePath string
