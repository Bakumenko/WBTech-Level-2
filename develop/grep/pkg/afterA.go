package pkg

import (
	"errors"
	"strconv"
	"strings"
)

func printLinesAfterSearched(textByLines []string, targerString string, countLines int, ignoreCase bool, lineNumber bool) ([]string, error) {
	if ignoreCase {
		targerString = strings.ToLower(targerString)
	}
	for index, line := range textByLines {
		if ignoreCase {
			line = strings.ToLower(line)
		}
		if strings.Contains(line, targerString) {
			lenSlice := min(len(textByLines)-index, countLines)
			result := make([]string, lenSlice)
			for i := index; i <= lenSlice; i++ {
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
