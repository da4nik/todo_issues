package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	configFileName = ".todo_issues.json"
)

// FileStruct - represents config file struct
type FileStruct struct {
	GithubAccessKey string `json:"github_access_key"`
}

var (
	defaultConfigFile string

	// GithubAccessKey - access key, which allows to create issues in repo
	GithubAccessKey string
)

// Load - loads config
func Load() {
	defaultConfigFile = os.Getenv("HOME") + "/" + configFileName

	loadFile()

	if os.Getenv("TI_GITHUB_ACCESS_KEY") != "" {
		GithubAccessKey = os.Getenv("TI_GITHUB_ACCESS_KEY")
	}
}

func loadFile() {
	read, err := ioutil.ReadFile(defaultConfigFile)
	if err != nil {
		fmt.Printf("Unable to load config file: %s (%s)\n", defaultConfigFile, err.Error())
		return
	}

	var data FileStruct
	if err := json.Unmarshal(read, &data); err != nil {
		fmt.Printf("Unable to parse config file: %s (%s)\n", defaultConfigFile, err.Error())
		return
	}

	GithubAccessKey = data.GithubAccessKey
}
