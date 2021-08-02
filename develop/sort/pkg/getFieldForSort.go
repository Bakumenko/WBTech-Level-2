package pkg

import "strings"

func getFieldForSortFromLines(numberColumn int, line string) string {
	if line == "" {
		return ""
	}

	var sortedField string
	if numberColumn < 0 {
		sortedField = line
	} else {
		lineBySpace := strings.Fields(line)
		if numberColumn < len(lineBySpace) {
			sortedField = lineBySpace[numberColumn]
		} else {
			sortedField = lineBySpace[0]
		}
	}
	return sortedField
}
