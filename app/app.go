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

func switchBranch(branchName string) {
	fmt.Println("Switching to branch:", branchName)
	err := exec.Command("git", "checkout", branchName).Run()
	if err != nil {
		err = exec.Command("git", "stash").Run()
		if err != nil {
			fmt.Println("Error stashing changes:", err)
		}
		err = exec.Command("git", "checkout", branchName).Run()
		if err != nil {
			fmt.Println("Error switching branches:", err)
		}
	}
}

func pullBranch() {
	err := exec.Command("git", "pull").Run()
	if err != nil {
		fmt.Println("Error pulling branch:", err)
	}
}

func mergeBranch(branchName string) {
	err := exec.Command("git", "switch", branchName).Run()
	if err != nil {
		fmt.Println("Error switching to branch:", err)
		return
	}
	err = exec.Command("git", "merge", branchName).Run()
	if err != nil {
		fmt.Println("Error merging branch:", err)
	}
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
