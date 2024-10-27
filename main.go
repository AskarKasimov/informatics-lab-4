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
	hasMinus := func(str string) bool {
		for index := range str {
			if str[index] == ' ' {
				continue
			}
			if str[index] == '-' {
				return true
			} else {
				break
			}
		}
		return false
	}
	excludeMinus := func(str string) string {
		for index := range str {
			if str[index] == ' ' || str[index] == '-' {
				continue
			}
			return str[index:]
		}
		panic("error")
	}
	countSpaces := func(str string) int {
		counter := 0
		for _, symbol := range str {
			if symbol == ' ' || symbol == '-' {
				counter += 1
			} else {
				break
			}
		}
		return counter / 2
	}
	makeJsonString := func(str string) string {
		if str[len(str)-1] == ':' {
			return "\"" + str[:len(str)-1] + "\":"
		}
		if !strings.Contains(str, ": ") {
			return "\"" + str + "\""
		}
		withColon := strings.Split(str, ": ")
		return "\"" + withColon[0] + "\":\"" + withColon[1] + "\""
	}
	countingOffset := func(slice []string) string {
		bracers := ""
		structure := ""
		previous := countSpaces(slice[0])
		for lineIndex := 1; lineIndex < len(slice); lineIndex++ {
			// if len(slice[lineIndex]) == 0 {
			// 	break
			// }
			current := countSpaces(slice[lineIndex])
			if previous < current {
				structure += makeJsonString(excludeMinus(slice[lineIndex-1]))
				structure += "["
				bracers += "["
				if hasMinus(slice[lineIndex]) {
					structure += "{"
					bracers += "{"
				}
			}
			if previous == current {
				if hasMinus(slice[lineIndex]) {
					structure += makeJsonString(excludeMinus(slice[lineIndex-1])) + "},{"
					// bracers = bracers[:len(bracers)-1] + "{"
				} else {
					structure += makeJsonString(excludeMinus(slice[lineIndex-1])) + ","
				}
			}
			if previous > current {
				// fmt.Println(previous,current, bracers, structure)
				structure += makeJsonString(excludeMinus(slice[lineIndex-1]))
				fmt.Println(bracers, len(bracers)-1, len(bracers)-(previous-current))
				for bracerIndex := len(bracers) - 1; bracerIndex >= len(bracers)-(previous-current); bracerIndex-- {
					if bracers[bracerIndex] == '{' {
						structure += "}"
					} else if bracers[bracerIndex] == '[' {
						structure += "]"
					}
					if bracerIndex != 0 && lineIndex != len(slice)-1 && bracers[bracerIndex] == '[' {
						structure += ","
					}
				}
				bracers = bracers[:len(bracers)-(previous-current)]
			}
			previous = current
		}
		return "{" + structure + "}"
	}
	content := strings.Split(readInput(), "\r\n")
	return countingOffset(content)
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
