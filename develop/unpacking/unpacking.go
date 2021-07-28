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

	for i := 0; i < len(sliceRunesString); i++ {
		val := sliceRunesString[i]
		stringVal := string(val)

		if stringVal == "\\" {
			if i + 1 < len(sliceRunesString) {
				next := sliceRunesString[i + 1]
				if string(next) == "\\" || unicode.IsDigit(next){
					current = next
					currentWritten = true
				} else {
					return "", errors.New("invalid string")
				}
			} else {
				return "", errors.New("invalid string")
			}
			i++
		} else if unicode.IsDigit(val) {
			if !currentWritten {
				return "", errors.New("invalid string")
			}
			k, err := strconv.Atoi(stringVal)
			if err != nil {
				return "", err
			}
			for j := 0; j < k; i++ {
				content = append(content, current)
				currentWritten = false
			}
		} else {
			if currentWritten {
				content = append(content, current)
			}
			current = val
			currentWritten = true
		}
	}

	return string(content), nil
}
