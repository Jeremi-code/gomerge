package main

import (
	"flag"
	"fmt"
	"github.com/jeremi-code/gomerge/app"
)

func main() {
	showRecent := flag.Bool("recent", false, "Open recent file")
	folderName := flag.String("folder", "", "Folder to open")
	branchName := flag.String("branch", "", "Branch to use")
	showRecentShort := flag.Bool("r", false, "Open recent file (short)")
	folderNameShort := flag.String("f", "", "Folder to open (short)")
	branchNameShort := flag.String("b", "", "Branch to use (short)")
	config := app.ReadConfigFile()
	flag.Parse()

	if *branchName != "" {
		config.BranchName = *branchName
		app.WriteConfigFile(config)
		fmt.Println("Branch is written")
	} else if *branchNameShort != "" {
		config.BranchName = *branchNameShort
		app.WriteConfigFile(config)
		fmt.Println("Branch is written")
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
		targetFolder = *folderName
	case flag.NArg() > 0:
		targetFolder = flag.Arg(0)
	default:
		fmt.Println("No file or folder specified, starting app...")
		return
	}

	app.Start(&config, targetFolder)
}
