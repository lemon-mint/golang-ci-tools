package main

import (
	"fmt"
	"log"

	"github.com/lemon-mint/golang-ci-tools/gocap"
	"github.com/lemon-mint/golang-ci-tools/staticcheck"
)

func main() {
	fmt.Println("=== gocap ===")
	// Run gocap.
	gocapInfo, err := gocap.Run(".")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", gocapInfo)

	fmt.Println("=== staticcheck ===")
	// Run staticcheck.
	staticcheckInfo, err := staticcheck.Run("./...")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", staticcheckInfo)
}
