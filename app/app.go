package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	BranchName string `json:"branch"`
	FolderDirectory string `json:"folder-directory"`
}

func Start(branch_name string) {
	
}

func RunCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func IsBackend() bool {
	fileName := GetFileName()
	var isBackend = strings.Contains(fileName, "api")

	return isBackend
}

func SwitchBranch(branchName string) {
	RunCommand("git stash")
	RunCommand("git checkout " + branchName)
}

func PullBranch() {
	RunCommand("git pull")
}

func MergeBranch(sourceBranchName string, targetBranchName string) {
	RunCommand("git switch " + targetBranchName)
	RunCommand("git merge " + sourceBranchName)
}

func GetBranch() (string, error) {
	config := ReadConfigFile()
	return config.BranchName, nil
}

func InitiateDocker() {
	OpenTerminal("docker compose up --build -w")
}

func UploadMigration() {
	InitiateDocker()
	OpenTerminal("yarn migrate")
}

func OpenTerminal(command string) {
	RunCommand("ghostty -c " + command)

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

func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func GetFileName() string {
	dir := GetCurrentDirectory()
	var dirSplit = strings.SplitAfter(dir, "/")
	var fileName = dirSplit[len(dirSplit)-1]
	if len(fileName) == 0 {
		log.Fatal("Error: File name is empty")
	}
	return fileName
}
