package pkg

import (
	"regexp"
	"strconv"
	"strings"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func grepLinesNearSearched(textByLines []string, targerString string, before int, after int, ignoreCase bool, lineNumber bool, regular bool, invert bool) ([]string, int, error) {
	var result []string
	var count int
	lineAlreadyResult := map[int]struct{}{}

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
				return nil, 0, err
			}
			if invert {
				matched = !matched
			}
			if matched {
				count++
				lenBefore := min(index, before)
				lenAfter := min(len(textByLines)-index-1, after)
				for i := index - lenBefore; i <= index+lenAfter; i++ {
					if lineNumber {
						textByLines[i] = strconv.Itoa(i+1) + ". " + textByLines[i]
					}

					if _, contain := lineAlreadyResult[i]; !contain {
						result = append(result, textByLines[i])
						lineAlreadyResult[i] = struct{}{}
					}
				}
			}
		} else {
			contain := strings.Contains(line, targerString)
			if invert {
				contain = !contain
			}
			if contain {
				count++
				lenBefore := min(index, before)
				lenAfter := min(len(textByLines)-index-1, after)
				for i := index - lenBefore; i <= index+lenAfter; i++ {
					if lineNumber {
						textByLines[i] = strconv.Itoa(i+1) + ". " + textByLines[i]
					}
					if _, alreadyIn := lineAlreadyResult[i]; !alreadyIn {
						result = append(result, textByLines[i])
						lineAlreadyResult[i] = struct{}{}
					}
				}
			}
		}
	}
	return result, count, nil
}
