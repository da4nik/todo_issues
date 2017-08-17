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

// fileStruct - represents config file struct
type fileStruct struct {
	GithubAccessKey string `json:"github_access_key"`
	Additions       bool   `json:"additions"`
}

var (
	defaultConfigFile string

	// GithubAccessKey - access key, which allows to create issues in repo
	GithubAccessKey string

	// Additions - whether or not add additional info
	Additions = false
)

// Load - loads config
func Load() {
	defaultConfigFile = os.Getenv("HOME") + "/" + configFileName

	loadFile()

	GithubAccessKey = getEnv("TI_GITHUB_ACCESS_KEY", GithubAccessKey)
}

func getEnv(name, value string) string {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	return env
}

func loadFile() {
	read, err := ioutil.ReadFile(defaultConfigFile)
	if err != nil {
		fmt.Printf("Unable to load config file: %s (%s)\n", defaultConfigFile, err.Error())
		return
	}

	var data fileStruct
	if err := json.Unmarshal(read, &data); err != nil {
		fmt.Printf("Unable to parse config file: %s (%s)\n", defaultConfigFile, err.Error())
		return
	}

	GithubAccessKey = data.GithubAccessKey
	Additions = data.Additions
}
