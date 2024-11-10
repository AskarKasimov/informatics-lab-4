package main

import (
	"fmt"
	"os"

	"github.com/AskarKasimov/informatics-lab-4/pkg"
)

func readInput() string {
	contentBytes, err := os.ReadFile("input.yaml")
	if err != nil {
		panic(err) // critical error: file not found
	}
	content := string(contentBytes)
	return content
}

func main() {
	fmt.Println(pkg.MainTask(readInput()))
	fmt.Println(pkg.FirstAdditionalTask(readInput()))
}
