package pkg

import (
	"errors"
	"strings"
)

func countOfSearchedInLines(textByLines []string, targerString string, ignoreCase bool) (int, error) {
	var count int
	if ignoreCase {
		targerString = strings.ToLower(targerString)
	}
	for _, line := range textByLines {
		if ignoreCase {
			line = strings.ToLower(line)
		}
		if strings.Contains(line, targerString) {
			count++
		}
	}
	if count == 0 {
		return 0, errors.New("substring didn't found")
	} else {
		return count, nil
	}
}
