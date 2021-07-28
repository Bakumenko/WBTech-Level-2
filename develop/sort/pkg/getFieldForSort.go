package pkg

import "strings"

func getFieldForSortFromLines(numberColumn int, line string) string {
	var sortedField string
	if numberColumn == -1 {
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
