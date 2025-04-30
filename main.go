package main

import (
	"os"
	"strings"
	"github.com/jeremi-code/gomerge/app"
)

func main() {

	if len(os.Args) == 2 {
		dir := app.GetCurrentDirectory()
		workFolder := app.ReadConfigFile().FolderDirectory
		if strings.Contains(dir, workFolder) {
			app.RunCommand("cd" + " " + os.Args[1])
			isBackend := app.IsBackend()
			if isBackend {
				app.UploadMigration()
				app.SwitchBranch("dev")
				app.MergeBranch("dev", "fix/issues")
			} else {
				app.SwitchBranch("div")
				app.MergeBranch("dev", "fix/issues")
			}

		}

	}
}
