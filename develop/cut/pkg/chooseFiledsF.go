package pkg

import (
	"sort"
	"strings"
)

func cutFields(lines []string, delimiter string, fields []int, outputLinesOnlyWithSeparator bool) []string {
	if delimiter == "" {
		return lines
	}

	var result []string
	sort.Ints(fields)
	for _, line := range lines {
		if strings.Contains(line, delimiter) {
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
		} else if !outputLinesOnlyWithSeparator {
			result = append(result, line)
		}
	}
	return result
}
