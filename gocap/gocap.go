package gocap

import (
	"bytes"
	"log"
	"os/exec"
)

func init() {
	// Set up the command.

	// go install github.com/cugu/gocap@latest
	cmd := exec.Command("go", "install", "github.com/cugu/gocap@latest")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Run(packageName string) ([]byte, error) {
	// Set up the command.
	cmd := exec.Command("gocap", "generate", packageName)
	// Set up the output buffer.
	var out bytes.Buffer
	cmd.Stdout = &out
	// Run the command.
	err := cmd.Run()
	// Return the output.
	return out.Bytes(), err
}
