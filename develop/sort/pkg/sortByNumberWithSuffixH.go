package pkg

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

var suffixes = map[string]int{
	"KB": 1 << 10,
	"MB": 1 << 20,
	"GB": 1 << 30,
	"TB": 1 << 40,
}

func parseStringWithSuffixInNumber(s string) (float64, error) {
	for key, val := range suffixes {
		i := strings.Index(s, key)
		if i > -1 {
			numberString := s[:i]
			number, err := strconv.ParseFloat(numberString, 64)
			if err != nil {
				return 0, err
			}
			numberInBytes := number * float64(val)
			return numberInBytes, nil
		}
	}
	return 0, errors.New("wrong number")
}

func sortByNumberWithSuffix(numberColumn int, lines []string, neededToReverse bool) ([]string, error) {
	var keys []float64
	keyColumnToLineMap := map[float64][]string{}
	for _, line := range lines {
		if line != "" {
			s := strings.Fields(line)
			if len(s) < numberColumn {
				return nil, errors.New("didn't find the column")
			}

			key, err := parseStringWithSuffixInNumber(s[numberColumn])
			if err != nil {
				return nil, err
			}
			keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
			keys = append(keys, key)
		}
	}

	if neededToReverse {
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
	} else {
		sort.Float64s(keys)
	}
	var result []string

	for _, key := range keys {
		result = append(result, keyColumnToLineMap[key]...)
	}

	return result, nil
}
