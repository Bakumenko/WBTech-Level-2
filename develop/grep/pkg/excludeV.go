package pkg

import (
	"strconv"
	"strings"
)

func excludeTargetStringFromLines(textByLines []string, targerString string, ignoreCase bool, lineNumber bool) ([]string, error) {
	result := []string{}
	if ignoreCase {
		targerString = strings.ToLower(targerString)
	}
	for index, line := range textByLines {
		if ignoreCase {
			line = strings.ToLower(line)
		}
		if !strings.Contains(line, targerString) {
			if lineNumber {
				line = strconv.Itoa(index+1) + ". " + line
			}
			result = append(result, line)
		}
	}
	return result, nil
}
