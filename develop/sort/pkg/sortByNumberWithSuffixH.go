package pkg

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

var suffixes = map[string]int{
	"KB": 1 << 10,
	"K":  1 << 10,
	"MB": 1 << 20,
	"M":  1 << 20,
	"GB": 1 << 30,
	"G":  1 << 30,
	"TB": 1 << 40,
	"T":  1 << 40,
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
	var result []string
	var notNumber []string
	for _, line := range lines {
		sortedField := getFieldForSortFromLines(numberColumn, line)
		key, err := parseStringWithSuffixInNumber(sortedField)
		if err != nil {
			notNumber = append(notNumber, line)
		} else {
			if _, ok := keyColumnToLineMap[key]; !ok {
				keys = append(keys, key)
			}
			keyColumnToLineMap[key] = append(keyColumnToLineMap[key], line)
		}
	}

	if neededToReverse {
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
		sort.Sort(sort.Reverse(sort.StringSlice(notNumber)))
		result = append(result, notNumber...)
	} else {
		result = append(result, notNumber...)
		for _, key := range keys {
			result = append(result, keyColumnToLineMap[key]...)
		}
	}
	return result, nil
}
