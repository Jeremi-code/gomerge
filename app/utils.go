package app

import (
	"log"
	"os"
	"strings"
)

func IsBackend() bool {
	fileName := GetFileName()
	var isBackend = strings.Contains(fileName, "api")
	return isBackend
}

func GetBranch() (string, error) {
	config := ReadConfigFile()
	return config.BranchName, nil
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

func RunWorkflow(branch string, isBackend bool) error {
    SwitchBranch("dev")
    PullBranch()
    MergeBranch("dev", branch)
    if isBackend {
        UploadMigration()
    }
    MergeBranch(branch, "dev")
    return nil
}