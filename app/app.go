package app

import (
	"log"
	"os"
	"path/filepath"
)

func Start(config *Config, targetFolder string) {
	workFolder := config.FolderDirectory
	branch := config.BranchName
	WriteConfigFile(*config)

	if branch == "" {
		branch = "fix-issues"
	}
	err := os.Chdir(filepath.Join(workFolder, targetFolder))
	if err != nil {
		log.Fatal("Error changing directory:", err)
		return
	}
	RunWorkflow(branch, targetFolder)
}
