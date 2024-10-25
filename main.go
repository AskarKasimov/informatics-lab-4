package main

import (
	"fmt"
	"os"
	"strings"

	"sigs.k8s.io/yaml"
)

func readInput() string {
	contentBytes, err := os.ReadFile("input.yaml")
	if err != nil {
		panic(err) // critical error: file not found
	}
	content := string(contentBytes)
	return content
}

func mainTask() string {
	content := strings.Split(strings.ReplaceAll(readInput(), " - ", "   "), "\r\n")
	for stringIndex := range content {
		if len(content[stringIndex]) == 0 {
			continue
		}
		if len(content[stringIndex])-strings.Index(content[stringIndex], ":") == 1 {
			content[stringIndex] = "\"" + content[stringIndex][:len(content[stringIndex])-1] + "\":{"
		} else {
			content[stringIndex] = "\"" + strings.Replace(content[stringIndex], ": ", "\":\"", 1) + "\","
		}
	}

	for _, str := range content {
		fmt.Println(str)
	}

	return "{" + strings.Join(content, "") + "}"
}

func firstAdditionalTask() string {
	content := readInput()
	res, err := yaml.YAMLToJSON([]byte(content))
	if err != nil {
		panic(err)
	}
	return string(res)
}

func main() {
	fmt.Println(mainTask())
}
