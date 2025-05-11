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
	dat, err := os.ReadFile(".myapprc")
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
	 fmt.Println("Writing config file with recent file:", config)

	dat, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dat))
	err = os.WriteFile("../.myapprc", dat, 0644)
	if err != nil {
		log.Fatal(err)
	}
}