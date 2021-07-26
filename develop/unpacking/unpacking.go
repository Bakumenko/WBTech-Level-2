package main

import (
	"errors"
	"unicode"
)

func UnpackingString(s string) (string, error) {
	r := []rune(s)
	content := []rune{}
	var last rune
	lastIsWritten := false

	for i := 0; i < len(r); i++ {
		val := r[i]
		if string(val) == "\\" {
			if i < len(r)-1 {
				if string(r[i+1]) == "\\" || unicode.IsDigit(r[i+1]) {
					last = r[i+1]
					lastIsWritten = true
					content = append(content, last)
					i++
				} else {
					return "", errors.New("invalid string")
				}
			} else {
				return "", errors.New("invalid string")
			}
		} else if unicode.IsDigit(val) {
			if lastIsWritten {
				for i := 0; i < int(val-'0')-1; i++ {
					content = append(content, last)
					lastIsWritten = false
				}
			} else {
				return "", errors.New("invalid string")
			}
		} else {
			content = append(content, val)
			last = val
			lastIsWritten = true
		}
	}

	return string(content), nil
}
