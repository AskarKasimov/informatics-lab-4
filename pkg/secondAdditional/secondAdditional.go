package secondAdditional

import (
	"regexp"
	"strings"
)

func excludeMinus(str string) string {
	re := regexp.MustCompile(`^[ -]*`)
	return re.ReplaceAllString(str, "")
}

func hasMinus(str string) bool {
	re := regexp.MustCompile("^(  )*(- )")
	return re.Match([]byte(str))
}

func countSpaces(str string) int {
	re := regexp.MustCompile("^(  )*(- )?")
	return len(re.FindString(str)) / 2
}

func makeJsonString(str string) string {
	reKeyOnly := regexp.MustCompile(`^(.+):$`)
	reKeyValue := regexp.MustCompile(`^(.+): (.+)$`)

	if reKeyOnly.MatchString(str) {
		// Если строка оканчивается на двоеточие, формируем ключ без значения
		key := reKeyOnly.FindStringSubmatch(str)[1]
		return "\"" + key + "\":"
	} else if reKeyValue.MatchString(str) {
		// Если строка содержит ключ и значение, разделенные ": "
		matches := reKeyValue.FindStringSubmatch(str)
		key, value := matches[1], matches[2]
		return "\"" + key + "\":\"" + value + "\""
	} else {
		// Если строка не содержит ": ", возвращаем ее как строку
		return "\"" + str + "\""
	}
}

func SecondAdditionalTask(input string) string {
	slice := strings.Split(input, "\r\n")
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
			if hasMinus(slice[lineIndex]) {
				structure += "[{"
				bracers += "[{"
			} else {
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
			// fmt.Println(bracers, len(bracers)-1, len(bracers)-(previous-current))
			for bracerIndex := len(bracers) - 1; bracerIndex >= len(bracers)-(previous-current)-1; bracerIndex-- {
				if bracerIndex >= len(bracers) || bracerIndex < 0 {
					continue
				}
				if bracers[bracerIndex] == '{' {
					structure += "}"
				} else if bracers[bracerIndex] == '[' {
					structure += "]"
				}
				// if bracerIndex != 0 && lineIndex != len(slice)-1 && bracers[bracerIndex] == '[' {
				// 	structure += ","
				// }
			}
			if structure[len(structure)-1] == '}' && lineIndex != len(slice)-1 {
				structure += ",{"
			}
			bracers = bracers[:len(bracers)-(previous-current)]
		}
		previous = current
	}
	return "{" + structure + "}"
}
