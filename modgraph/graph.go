package modgraph

import (
	"bytes"
	"os/exec"
	"strings"
)

func Run() []string {
	// go mod graph
	cmd := exec.Command("go", "mod", "graph")
	// Set up the output buffer.
	var out bytes.Buffer
	cmd.Stdout = &out
	// Run the command.
	cmd.Run()
	// Return the output.
	output := out.String()
	splited := strings.Split(output, "\n")
	// remove the empty line
	cleaned := make([]string, 0, len(splited))
	for i := 0; i < len(splited); i++ {
		if len(splited[i]) > 0 {
			cleaned = append(cleaned, splited[i])
		}
	}
	return cleaned
}
