package pkg

import (
	"sort"
	"strings"
)

func simpleSort(lines []string, neededToReverse bool) []string {
	var lowerLinesSlice []string
	mapLowerLinesToLines := map[string]string{}

	for _, line := range lines {
		lowerLine := strings.ToLower(line)
		lowerLinesSlice = append(lowerLinesSlice, lowerLine)
		mapLowerLinesToLines[lowerLine] = line
	}

	if neededToReverse {
		sort.Sort(sort.Reverse(sort.StringSlice(lowerLinesSlice)))
	} else {
		sort.Strings(lowerLinesSlice)
	}
	var result []string
	for _, key := range lowerLinesSlice {
		result = append(result, mapLowerLinesToLines[key])
	}
	return result
}
