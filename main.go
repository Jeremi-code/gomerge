package main

import (
	// "flag"
	"fmt"
	"github.com/jeremi-code/gomerge/app"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		branchName,err:=app.GetBranch()
		if(err!=nil){
			panic(err)
		}
		fmt.Println("Current branch name:", branchName)
		
	}
	fmt.Println("Starting the app...", os.Args[0])
	// branch := flag.String("branch", "", "Branch name to start the app with")
	// flag.Parse()

	// app.Start(*branch)

}
