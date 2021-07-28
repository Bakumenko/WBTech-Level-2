package main

import (
	"errors"
	"strconv"
	"unicode"
)

func UnpackingString(s string) (string, error) {
	content := []rune{}
	sliceRunesString := []rune(s)
	var current rune
	var currentWritten bool
	i := 0
	for i < len(sliceRunesString) {
		val := sliceRunesString[i]
		stringVal := string(val)

		if stringVal == "\\" {
			if i+1 < len(sliceRunesString) {
				next := sliceRunesString[i+1]
				if string(next) == "\\" || unicode.IsDigit(next) {
					if currentWritten {
						content = append(content, current)
					}
					current = next
					currentWritten = true
				} else {
					return "", errors.New("invalid string")
				}
			} else {
				return "", errors.New("invalid string")
			}
			i++
			i++
		} else if unicode.IsDigit(val) {
			if !currentWritten {
				return "", errors.New("invalid string")
			}
			var num int
			for unicode.IsDigit(sliceRunesString[i]) {
				k, err := strconv.Atoi(string(sliceRunesString[i]))
				if err != nil {
					return "", err
				}
				num = num*10 + k
				i++
				if i >= len(sliceRunesString) {
					break
				}
			}
			for j := 0; j < num; j++ {
				content = append(content, current)
			}
			currentWritten = false
		} else {
			if currentWritten {
				content = append(content, current)
			}
			current = val
			currentWritten = true
			i++
		}
	}
	if currentWritten {
		content = append(content, current)
	}
	return string(content), nil
}
