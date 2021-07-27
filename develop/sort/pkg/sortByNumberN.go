package pkg

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

func sortByNumber(numberColumn int, lines []string, neededToReverse bool) ([]string, error) {
	var keys []int
	keyColumnToLineMap := map[int][]string{}
	for _, line := range lines {
		if line != "" {
			s := strings.Fields(line)
			if len(s) < numberColumn {
				return nil, errors.New("didn't find the column")
			}

			key, err := strconv.Atoi(s[numberColumn])
			if err != nil {
				return nil, err
			}
			keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
			keys = append(keys, key)
		}
	}

	if neededToReverse {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	} else {
		sort.Ints(keys)
	}
	var result []string

	for _, key := range keys {
		result = append(result, keyColumnToLineMap[key]...)
	}

	return result, nil
}
