package gosec

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/acarl005/stripansi"
)

func init() {
	// Set up the command.

	// go install github.com/securego/gosec/v2/cmd/gosec@latest
	cmd := exec.Command("go", "install", "github.com/securego/gosec/v2/cmd/gosec@latest")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Run(packageName string) []byte {
	cmd := exec.Command("gosec", packageName)
	// Set up the output buffer.
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	// Run the command.
	err := cmd.Run()
	if err != nil {
		log.Println("Error running gosec:", err)
	}
	// Return the output.
	return []byte(stripansi.Strip(out.String()))
}
