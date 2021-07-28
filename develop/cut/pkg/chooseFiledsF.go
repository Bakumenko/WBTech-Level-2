package pkg

import (
	"sort"
	"strings"
)

func cutFields(lines []string, delimiter string, fields []int, outputLinesOnlyWithSeparator bool) []string {
	var result []string
	sort.Ints(fields)
	for _, line := range lines {
		if !strings.Contains(line, delimiter) && !outputLinesOnlyWithSeparator {
			println()
			result = append(result, line)
		} else {
			lineSliceByDelimiter := strings.Split(line, delimiter)
			var newLine string
			for _, val := range fields {
				val--
				if val < len(lineSliceByDelimiter) {
					if newLine != "" {
						newLine += delimiter + lineSliceByDelimiter[val]
					} else {
						newLine += lineSliceByDelimiter[val]
					}
				}
			}
			result = append(result, newLine)
		}
	}
	return result
}
