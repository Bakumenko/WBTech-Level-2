package pkg

import (
	"errors"
	"strconv"
	"strings"
)

func printLinesNearSearched(textByLines []string, targerString string, countLines int, ignoreCase bool, lineNumber bool) ([]string, error) {
	if ignoreCase {
		targerString = strings.ToLower(targerString)
	}
	for index, line := range textByLines {
		if ignoreCase {
			line = strings.ToLower(line)
		}
		if strings.Contains(line, targerString) {
			lenBefore := min(index, countLines)
			lenAfter := min(len(textByLines)-index, countLines)
			lenSlice := lenBefore + lenAfter
			result := make([]string, lenSlice)
			for i := index - lenBefore; i <= index+lenAfter; i++ {
				if lineNumber {
					textByLines[i] = strconv.Itoa(i+1) + ". " + textByLines[i]
				}
				result = append(result, textByLines[i])
			}
			return result, nil
		}
	}
	return nil, errors.New("substring didn't found")
}
