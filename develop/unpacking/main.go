package main

import (
	"fmt"
)

func main() {
	slice := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	for _, s := range slice {
		res, err := UnpackingString(s)
		if err != nil {
			fmt.Printf("%v => `\" \" (%v)\n", s, err)
		} else {
			fmt.Printf("%v => %v\n", s, res)
		}
	}
}
