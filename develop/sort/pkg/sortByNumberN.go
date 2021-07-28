package pkg

import (
	"sort"
	"strconv"
)

func sortByNumber(numberColumn int, lines []string, neededToReverse bool) ([]string, error) {
	if numberColumn == -1 {
		numberColumn = 0
	}
	var keys []int
	var result []string
	var notNumber []string
	keyColumnToLineMap := map[int][]string{}
	for _, line := range lines {
		sortedField := getFieldForSortFromLines(numberColumn, line)
		key, err := strconv.Atoi(sortedField)
		if err != nil {
			notNumber = append(notNumber, line)
		} else {
			if _, ok := keyColumnToLineMap[key]; !ok {
				keys = append(keys, key)
			}
			keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
		}
	}

	sort.Strings(notNumber)
	if neededToReverse {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
		sort.Sort(sort.Reverse(sort.StringSlice(notNumber)))
		result = append(result, notNumber...)
	} else {
		result = append(result, notNumber...)
		sort.Ints(keys)
		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
	}
	return result, nil
}
