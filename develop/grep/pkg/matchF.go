package pkg

import (
	"errors"
	"strconv"
	"strings"
)

func fullMatchSearchedWithLine(textByLines []string, targerString string, countLines int, ignoreCase bool, lineNumber bool) ([]string, error) {
	var result []string

	if ignoreCase {
		targerString = strings.ToLower(targerString)
	}
	for index, line := range textByLines {
		if ignoreCase {
			line = strings.ToLower(line)
		}

		if lineNumber {
			line = strconv.Itoa(index+1) + ". " + line
		}

		if strings.Contains(line, targerString) {
			result = append(result, line)
		}
	}
	if result == nil {
		return nil, errors.New("substring didn't found")
	} else {
		return result, nil
	}
}
