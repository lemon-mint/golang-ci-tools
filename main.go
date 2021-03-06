package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/lemon-mint/golang-ci-tools/gocap"
	"github.com/lemon-mint/golang-ci-tools/golicenses"
	"github.com/lemon-mint/golang-ci-tools/gosec"
	"github.com/lemon-mint/golang-ci-tools/modgraph"
	"github.com/lemon-mint/golang-ci-tools/staticcheck"
)

const Version = "0.1.2-dev"

func main() {
	var pkgname string = "."
	if len(os.Args) > 1 {
		pkgname = os.Args[1]
	}
	fmt.Print("# Golang CI Tools Report\n\n")
	fmt.Print("Report generated at: ", time.Now().Format(time.RFC3339))
	fmt.Print("\n\n")
	fmt.Print("Go version: ", runtime.Version())
	fmt.Print("\n\n")
	fmt.Print("golang-ci-tools version: ", Version)
	fmt.Print("\n\n")

	fmt.Print("## staticcheck\n\n")
	fmt.Print("```\n")
	// Run staticcheck.
	staticcheckInfo := staticcheck.Run("./...")
	if len(staticcheckInfo) > 0 {
		fmt.Printf("%s\n", staticcheckInfo)
	} else {
		fmt.Println("🎉  No staticcheck errors found!")
	}
	fmt.Print("```\n\n")

	fmt.Print("## gosec\n\n")
	fmt.Print("```\n")
	// Run gosec.
	gosecInfo, err := gosec.Run("./...")
	if err != nil {
		fmt.Print("** Issue Detected **\n\n")
	}
	if len(gosecInfo) > 0 {
		fmt.Printf("%s\n", gosecInfo)
	}
	fmt.Print("```\n\n")

	fmt.Print("## gocap\n\n")
	fmt.Print("```\n")
	// Run gocap.
	gocapInfo, err := gocap.Run(pkgname)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s\n", gocapInfo)
	fmt.Print("```\n\n")

	fmt.Print("## go-licenses\n\n")
	// Run go-licenses.
	golicensesInfo, err := golicenses.Run("./...")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s\n", golicensesInfo)
	fmt.Print("\n\n")

	fmt.Print("## Dependencies\n\n")
	deps := modgraph.Run()
	fmt.Printf("Total dependencies: %d\n", len(deps))
	fmt.Printf("<details><summary>Show Full Dependencies</summary>\n\n")
	for _, dep := range deps {
		fmt.Printf(" - %s\n", dep)
	}
	fmt.Printf("</details>\n\n")
}
