package app

import (
	"fmt"
	"log"
	"os/exec"
)

func RunCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	fmt.Println("Running command:", command)
	if err != nil {
		fmt.Println("Error executing command:", command)
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func OpenTerminal(command string) {
	fmt.Println("Opening terminal with command:", command)
	cmd := exec.Command("ghostty", "-e", "bash", "-c", "\""+command+"\"")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to start Ghostty: %v", err)
	}
}
