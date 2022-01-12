package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"os/exec"
)

func init() {
	// Set up the command.

	// go install github.com/google/go-licenses@latest
	cmd := exec.Command("go", "install", "github.com/google/go-licenses@latest")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Run(packageName string) ([]byte, error) {
	cmd := exec.Command("go-licenses", "csv", packageName)
	// Set up the output buffer.
	var out bytes.Buffer
	cmd.Stdout = &out
	// Run the command.
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	// Return the output.
	records, err := csv.NewReader(&out).ReadAll()
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
