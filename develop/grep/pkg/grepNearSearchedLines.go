package pkg

import (
	"regexp"
	"strconv"
	"strings"
)

func grepLinesNearSearched(textByLines []string, targerString string, before int, after int, ignoreCase bool, lineNumber bool, regular bool) ([]string, error) {
	var result []string
	if ignoreCase {
		targerString = strings.ToLower(targerString)
	}
	for index, line := range textByLines {
		if ignoreCase {
			line = strings.ToLower(line)
		}
		if regular {
			matched, err := regexp.MatchString(targerString, line)
			if err != nil {
				return nil, err
			}
			if matched {
				if lineNumber {
					line = strconv.Itoa(index+1) + ". " + line
				}

				result = append(result, line)
			}
		} else {
			if strings.Contains(line, targerString) {
				lenBefore := min(index, before)
				lenAfter := min(len(textByLines)-index-1, after)
				for i := index - lenBefore; i <= index+lenAfter; i++ {
					if lineNumber {
						textByLines[i] = strconv.Itoa(i+1) + ". " + textByLines[i]
					}
					result = append(result, textByLines[i])
				}
			}
		}
	}
	return result, nil
}
