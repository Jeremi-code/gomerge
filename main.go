package main

import (
	"flag"
	"fmt"
	"github.com/jeremi-code/gomerge/app"
)

func main() {
	showRecent := flag.Bool("recent", false, "Open recent file")
	folderName := flag.String("folder", "", "Folder to open")
	config := app.ReadConfigFile()
	flag.Parse()

	var targetFolder string

	switch {
	case *showRecent:
		if config.RecentFolder == "" {
			fmt.Println("No recent folder found.", config)
			return
		}
		fmt.Println("what is then")
		targetFolder = config.RecentFolder
	case *folderName != "":
		targetFolder = *folderName
	case flag.NArg() > 0:
		targetFolder = flag.Arg(0)
	default:
		fmt.Println("No file or folder specified, starting app...")
		return
	}

	app.Start(&config, targetFolder)
}
