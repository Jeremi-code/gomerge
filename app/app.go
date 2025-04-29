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
}

func Start(branch_name string) {
	var isBackend bool = isBackend()
	if isBackend {
		fmt.Println("this is backend project")
	}
	fmt.Println("Starting the app with branch name:", branch_name)
}

func RunCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func isBackend() bool {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	var dirSplit = strings.SplitAfter(dir, "/")
	var fileName = dirSplit[len(dirSplit)-1]
	if len(fileName) == 0 {
		log.Fatal("Error: File name is empty")
	}
	var isBackend = strings.Contains(fileName, "api")

	return isBackend
}

func SwitchBranch(branchName string) {
	fmt.Println("Switching to branch:", branchName)
	RunCommand("git checkout " + branchName)
	RunCommand("git stash")
	RunCommand("git checkout " + branchName)
}

func PullBranch() {
	RunCommand("git pull")
}

func MergeBranch(branchName string) {
	RunCommand("git switch " + branchName)
	RunCommand("git merge " + branchName)
}

func GetBranch() (string, error) {
	dat, err := os.ReadFile(".myapprc")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}
	var config Config
	err = json.Unmarshal(dat, &config)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return "", err
	}
	return config.BranchName, nil
}

func InitiateDocker() {
	RunCommand("docker compose up --build -w")
}

func UploadMigration() {
	InitiateDocker()
	RunCommand("yarn")
}
