package pkg

import (
	"strconv"
	"strings"
)

func grepLinesNearSearched(textByLines []string, targerString string, before int, after int, ignoreCase bool, lineNumber bool) ([]string, error) {
	var result []string
	if ignoreCase {
		targerString = strings.ToLower(targerString)
	}
	for index, line := range textByLines {
		if ignoreCase {
			line = strings.ToLower(line)
		}
		if strings.Contains(line, targerString) {
			lenBefore := min(index, before)
			lenAfter := min(len(textByLines)-index-1, after)
			//lenSlice := lenBefore + lenAfter
			for i := index - lenBefore; i <= index+lenAfter; i++ {
				if lineNumber {
					textByLines[i] = strconv.Itoa(i+1) + ". " + textByLines[i]
				}
				result = append(result, textByLines[i])
			}
		}
	}
	return result, nil
}
