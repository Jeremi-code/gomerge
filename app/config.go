package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	BranchName      string `json:"branch"`
	FolderDirectory string `json:"folder-directory"`
	RecentFile      string `json:"recent-file"`
}

func ReadConfigFile() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dat, err := os.ReadFile(homeDir + "/.visrc")
	if err != nil {
		log.Fatal(err)
	}
	var config Config

	err = json.Unmarshal(dat, &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config file read successfully", string(dat))
	return config
}

func WriteConfigFile(recentFile string) {
	config := ReadConfigFile()
	config.RecentFile = recentFile

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Writing config file with recent file:", config)

	dat, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dat))
	err = os.WriteFile(homeDir+"/.visrc", dat, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
