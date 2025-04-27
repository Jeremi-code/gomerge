package main

import (
	"flag"
	"github.com/jeremi-code/gomerge/app"
)

func main() {
	branch := flag.String("branch", "", "Branch name to start the app with")
	flag.Parse()

	app.Start(*branch)

}
