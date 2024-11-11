package main

import (
	"fmt"
	"os"

	"github.com/AskarKasimov/informatics-lab-4/pkg/firstAdditional"
	"github.com/AskarKasimov/informatics-lab-4/pkg/mainTask"
	"github.com/AskarKasimov/informatics-lab-4/pkg/secondAdditional"
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
	fmt.Println(mainTask.MainTask(readInput()))
	fmt.Println(firstAdditional.FirstAdditionalTask(readInput()))
	fmt.Println(secondAdditional.SecondAdditionalTask(readInput()))
}
