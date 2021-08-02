package pkg

import (
	"regexp"
	"strconv"
	"strings"
)

func grepExcludeTargetStringFromLines(textByLines []string, targerString string, ignoreCase bool, lineNumber bool, regular bool) ([]string, error) {
	result := []string{}
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
			if !matched {
				if lineNumber {
					line = strconv.Itoa(index+1) + ". " + line
				}

				result = append(result, line)
			}
		} else {
			if !strings.Contains(line, targerString) {
				if lineNumber {
					line = strconv.Itoa(index+1) + ". " + line
				}

				result = append(result, line)
			}
		}
	}
	return result, nil
}
