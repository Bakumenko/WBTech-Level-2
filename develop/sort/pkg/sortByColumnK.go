package pkg

import (
	"errors"
	"sort"
	"strings"
)

func sortByColumn(numberColumn int, lines []string) ([]string, error) {
	var keys []string
	keyColumnToLineMap := map[string][]string{}
	for _, line := range lines {
		s := strings.Fields(line)
		if len(s) < numberColumn {
			return nil, errors.New("didn't find the column")
		}

		if line != "" {
			key := s[numberColumn]
			keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)
	var result []string

	for _, key := range keys {
		result = append(result, keyColumnToLineMap[key]...)
	}

	return result, nil
}
