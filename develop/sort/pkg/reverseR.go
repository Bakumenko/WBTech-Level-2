package pkg

import (
	"sort"
)

func reverseSort(lines []string) ([]string, error) {
	sort.Strings(lines)
	for i := 0; i < len(lines)/2; i++ {
		indexOfEndSlice := len(lines) - 1 - i
		lines[i], lines[indexOfEndSlice] = lines[indexOfEndSlice], lines[i]
	}

	return lines, nil
}
