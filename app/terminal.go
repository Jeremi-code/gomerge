package app

import (
	"os/exec"
	"fmt"
	"log"
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
	cmd := exec.Command("ghostty", "-e", command)
    err := cmd.Start()
    if err != nil {
        log.Fatalf("Failed to start Ghostty: %v", err)
    }

}