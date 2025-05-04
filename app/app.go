package app

import (
	"os"
)

func Start() {
	dir := GetCurrentDirectory()
	workFolder := ReadConfigFile().FolderDirectory
	branch, err := GetBranch()
	if err != nil {
		return
	}

	if branch == "" {
		branch = "fix-issues"
	}

	if dir == workFolder {
		os.Chdir(os.Args[1])
		RunWorkflow(branch, IsBackend())
	} else {
		err := os.Chdir(workFolder + "/" + os.Args[1])
		if err != nil {
			return
		}
		RunWorkflow(branch, IsBackend())
	}
}
