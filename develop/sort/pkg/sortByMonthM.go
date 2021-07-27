package pkg

import (
	"errors"
	"sort"
	"strings"
)

type Month []string

// map to store weekdays' relative order
var month = map[string]int{
	"january":   1,
	"fabruary":  2,
	"march":     3,
	"april":     4,
	"may":       5,
	"june":      6,
	"july":      7,
	"august":    8,
	"september": 9,
	"october":   10,
	"november":  11,
	"december":  12,
}

func (m Month) Len() int           { return len(m) }
func (m Month) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Month) Less(i, j int) bool { return month[m[i]] < month[m[j]] }

func sortByMonth(numberColumn int, lines []string, neededToReverse bool) ([]string, error) {
	var keys Month
	keyColumnToLineMap := map[string][]string{}
	for _, line := range lines {
		if line != "" {
			s := strings.Fields(line)
			if len(s) < numberColumn {
				return nil, errors.New("didn't find the column")
			}

			if numberColumn < 0 {
				return nil, errors.New("error number of column")
			}

			key := s[numberColumn]
			keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
			keys = append(keys, key)
		}
	}

	if neededToReverse {
		sort.Sort(sort.Reverse(keys))
	} else {
		sort.Sort(keys)
	}
	var result []string

	for _, key := range keys {
		result = append(result, keyColumnToLineMap[key]...)
	}

	return result, nil
}
