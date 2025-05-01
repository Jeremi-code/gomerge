package app

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	BranchName      string `json:"branch"`
	FolderDirectory string `json:"folder-directory"`
}

func ReadConfigFile() Config {
	dat, err := os.ReadFile(".myapprc")
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = json.Unmarshal(dat, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
