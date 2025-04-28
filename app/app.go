package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

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
