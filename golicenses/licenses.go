package golicenses

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/lemon-mint/golang-ci-tools/markdown/table"
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
	tablerecords := make([][]string, len(records))
	for i, record := range records {
		tablerecords[i] = make([]string, len(record))
		tablerecords[i][0] = fmt.Sprintf("[%s](https://pkg.go.dev/%s)", record[0], record[0])
		tablerecords[i][1] = fmt.Sprintf("[%s](%s)", record[1], record[1])
		tablerecords[i][2] = record[2]
	}

	gfm := table.NewTable([]string{"Package Name", "License File", "License"}, tablerecords)
	return gfm, nil
}
