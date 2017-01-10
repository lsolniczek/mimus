package fileUtil

import "io/ioutil"

func ProjectExist(path string, name string) bool {
	for _, pn := range projectsList(path) {
		if pn == name {
			return true
		}
	}
	return false
}

func projectsList(path string) []string {
	var p []string
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return p
	}
	for _, info := range fileInfo {
		if info.IsDir() {
			p = append(p, info.Name())
		}
	}
	return p
}
