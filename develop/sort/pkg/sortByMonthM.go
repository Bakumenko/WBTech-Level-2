package pkg

import (
	"errors"
	"sort"
	"strings"
)

type Month []string

// map to store weekdays' relative order
var month = map[string]int{
	"January":   1,
	"Fabruary":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func (m Month) Len() int           { return len(m) }
func (m Month) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Month) Less(i, j int) bool { return month[m[i]] < month[m[j]] }

func sortByMonth(numberColumn int, lines []string) ([]string, error) {
	var keys Month
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

	sort.Sort(keys)
	var result []string

	for _, key := range keys {
		result = append(result, keyColumnToLineMap[key]...)
	}

	return result, nil
}
