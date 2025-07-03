package app

import (
	"strings"
)

func IsBackend(folderName string) bool {
	var isBackend = strings.Contains(folderName, "api")
	return isBackend
}

// func GetBranch() (string, error) {
// 	config := ReadConfigFile()
// 	return config.BranchName, nil
// }

// func GetCurrentDirectory() string {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return dir
// }

// func GetFileName() string {
// 	dir := GetCurrentDirectory()
// 	var dirSplit = strings.SplitAfter(dir, "/")
// 	var fileName = dirSplit[len(dirSplit)-1]
// 	if len(fileName) == 0 {
// 		log.Fatal("Error: File name is empty")
// 	}
// 	return fileName
// }

func RunWorkflow(branch string,folderName string) error {
    SwitchBranch("dev")
    PullBranch()
    MergeBranch("dev", branch)
    if IsBackend(folderName) {
        UploadMigration()
    }
    MergeBranch(branch, "dev")
    return nil
}