package app

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	BranchName      string `json:"branch"`
	FolderDirectory string `json:"folder-directory"`
	RecentFolder    string `json:"recent-folder"`
}

func ReadConfigFile() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to read config file", err)
	}
	dat, err := os.ReadFile(homeDir + "/.visrc")
	if err != nil {
		log.Fatal("Unable to read config file", err)
	}
	var config Config

	err = json.Unmarshal(dat, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func SetBranch(branchName string) Config {
	config := ReadConfigFile()
	config.BranchName = branchName
	return config
}

func WriteConfigFile(config Config) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Unable to write into config file", err)
	}

	dat, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal("Unable to marshal config to JSON", err)
	}
	
	err = os.WriteFile(homeDir+"/.visrc", dat, 0644)
	if err != nil {
		log.Fatal("Unable to write into config file", err)
	}
}
