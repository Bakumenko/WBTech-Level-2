package pkg

import (
	"sort"
	"strings"
)

func sortByColumn(numberColumn int, lines []string, neededToReverse bool) []string {
	keys := []string{}
	keyColumnToLineMap := map[string][]string{}
	for _, line := range lines {
		sortedField := getFieldForSortFromLines(numberColumn, line)
		key := strings.ToLower(sortedField)
		if _, ok := keyColumnToLineMap[key]; !ok {
			keys = append(keys, key)
		}
		keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
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
	return result
}
