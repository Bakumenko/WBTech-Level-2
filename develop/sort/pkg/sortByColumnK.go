package pkg

import (
	"sort"
	"strings"
)

func sortByColumn(numberColumns []int, lines []string, neededToReverse bool) ([]string, error) {
	var keys []string
	keyColumnToLineMap := map[string][]string{}
	for _, line := range lines {
		if line != "" {
			lineBySpace := strings.Fields(line)
			var lineByNeededColumns []string
			for _, val := range numberColumns {
				if val < len(lineBySpace) {
					lineByNeededColumns = append(lineByNeededColumns, lineBySpace[val])
				}
			}

			key := strings.Join(lineByNeededColumns, " ")
			keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
			keys = append(keys, key)
		}
	}
	if neededToReverse {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}
	var result []string

	for _, key := range keys {
		result = append(result, keyColumnToLineMap[key]...)
	}

	return result, nil
}
