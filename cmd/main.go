package main

import (
	"fmt"
	"os"
	"time"

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
	fmt.Println("Введите номер задания (0 - обязательное, 1, 2, 4): ")
	var task int
	fmt.Scan(&task)
	switch task {
	case 0:
		fmt.Println(mainTask.MainTask(readInput()))
	case 1:
		fmt.Println(firstAdditional.FirstAdditionalTask(readInput()))
	case 2:
		fmt.Println(secondAdditional.SecondAdditionalTask(readInput()))
	case 4:
		start0 := time.Now()
		fmt.Println(mainTask.MainTask(readInput()))
		end0 := time.Since(start0) * 100

		start1 := time.Now()
		fmt.Println(firstAdditional.FirstAdditionalTask(readInput()))
		end1 := time.Since(start1) * 100

		start2 := time.Now()
		fmt.Println(secondAdditional.SecondAdditionalTask(readInput()))
		end2 := time.Since(start2) * 100

		fmt.Println("Время выполнения базового скрипта: ", end0)
		fmt.Println("Время выполнения скрипта из библиотеки: ", end1)
		fmt.Println("Время выполнения улучшенного базового скрипта с регулярками: ", end2)
	}
}
