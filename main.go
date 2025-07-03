package main

import (
	"flag"
	"fmt"
	"github.com/jeremi-code/gomerge/app"
)

func main() {
	showRecent := flag.Bool("recent", false, "Open recent file")
	showRecentShort := flag.Bool("r", false, "Open recent file (short)")

	folderName := flag.String("folder", "", "Folder to open")
	folderNameShort := flag.String("f", "", "Folder to open (short)")

	branchName := flag.String("branch", "", "Branch to use")
	branchNameShort := flag.String("b", "", "Branch to use (short)")

	flag.Parse()

	config := app.ReadConfigFile()

	if *branchNameShort != "" {
		config.BranchName = *branchNameShort
		fmt.Println("Branch is written")
		app.WriteConfigFile(config)
	} else if *branchName != "" {
		config.BranchName = *branchName
		fmt.Println("Branch is written")
		app.WriteConfigFile(config)
	}

	var targetFolder string

	switch {
	case *showRecent || *showRecentShort:
		if config.RecentFolder == "" {
			fmt.Println("No recent folder found.", config)
			return
		}
		targetFolder = config.RecentFolder

	case *folderName != "" || *folderNameShort != "":
		if *folderNameShort != "" {
			targetFolder = *folderNameShort
		} else {
			targetFolder = *folderName
		}

	case flag.NArg() > 0:
		targetFolder = flag.Arg(0)

	default:
		fmt.Println("No file or folder specified, starting app...")
		return
	}

	app.Start(&config, targetFolder)
}
