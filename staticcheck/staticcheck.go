package staticcheck

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func init() {
	// Set up the command.

	// go install honnef.co/go/tools/cmd/staticcheck@latest
	cmd := exec.Command("go", "install", "honnef.co/go/tools/cmd/staticcheck@latest")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Run(packageName string) ([]byte, error) {
	cmd := exec.Command("staticcheck", packageName)
	// Set up the output buffer.
	var out bytes.Buffer
	cmd.Stdout = &out
	// Run the command.
	err := cmd.Run()
	// Return the output.
	return out.Bytes(), err
}
