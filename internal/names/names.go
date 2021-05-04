package names

import (
	"regexp"
	"strings"
)

func PathToPascalCase(path string) string {
	arr := regexp.MustCompile(`[^a-zA-Z]+`).Split(path, -1)

	var result string
	for _, el := range arr {
		result += Capitalize(el)
	}

	return result
}

func PathToSnakeCase(path string) string {
	arr := regexp.MustCompile(`[^a-zA-Z]+`).Split(path, -1)

	var result []string
	for _, el := range arr {
		if len(el) == 0 {
			continue
		}

		result = append(result, strings.ToLower(el))
	}

	return strings.Join(result, "_")
}

func Capitalize(str string) string {
	switch len(str) {
	case 0:
		return ""
	default:
		return strings.ToUpper(string(str[0])) + str[1:]
	}
}
