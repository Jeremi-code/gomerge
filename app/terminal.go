package app

import (
	"os/exec"
	"fmt"
	"log"
)

func RunCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func OpenTerminal(command string) {
	RunCommand("ghostty -c " + command)

}