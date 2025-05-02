package app

import (
	"os"
	"log"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func GetFileName() string {
	dir := GetCurrentDirectory()
	var dirSplit = strings.SplitAfter(dir, "/")
	var fileName = dirSplit[len(dirSplit)-1]
	if len(fileName) == 0 {
		log.Fatal("Error: File name is empty")
	}
	return fileName
}