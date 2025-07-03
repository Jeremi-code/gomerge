package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/jeremi-code/gomerge/app"
)

func main() {
	showRecent := flag.Bool("recent", false, "Open recent file")
	showRecentShort := flag.Bool("r", false, "Open recent file (short)")

	folderName := flag.String("folder", "", "Folder to open")
	folderNameShort := flag.String("f", "", "Folder to open (short)")

	branchName := flag.String("branch", "", "Branch to use")
	branchNameShort := flag.String("b", "", "Branch to use (short)")

	help := flag.Bool("help", false, "Show help")
	helpShort := flag.Bool("h", false, "Show help (short)")

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
	case *help || *helpShort:
		color.Cyan(`
GoMerge - A modern merge tool for vistec

Flags:
  -recent, -r           Open recent file
  -folder, -f <name>    Folder to open
  -branch, -b <name>    Branch to use
  -help, -h             Show this help message

Examples:
  gomerge -recent
  gomerge -folder myfolder
  gomerge -branch feature-xyz
`)

	case flag.NArg() > 0:
		targetFolder = flag.Arg(0)

	default:
		color.Green(`
  ________            _____                                  
 /  _____/   ____    /     \    ____ _______   ____    ____  
/   \  ___  /  _ \  /  \ /  \ _/ __ \\_  __ \ / ___\ _/ __ \ 
\    \_\  \(  <_> )/    Y    \\  ___/ |  | \// /_/  >\  ___/ 
 \______  / \____/ \____|__  / \___  >|__|   \___  /  \___  >
        \/                 \/      \/       /_____/       \/ 
`)
		color.Cyan("Welcome to GoMerge!")
		color.Green("Ready to merge your branches with style.")
		return
	}

	app.Start(&config, targetFolder)
}
