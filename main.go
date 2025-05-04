package main

import (
	"fmt"
	"os"

	"github.com/jeremi-code/gomerge/app"
)

func main() {
	fmt.Println("Welcome to GoMerge")
	if len(os.Args) == 2 {
		app.Start()
	}
}
