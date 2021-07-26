package main

import (
	"fmt"
)

func main() {
	slice := []string{
		"пятак",
		"листок",
		"олег",
		"пятка",
		"тяпка",
		"ЛЕГО",
		"слиток",
		"мукА",
		"столик",
		"акума",
	}

	result := SearchAnagramByDictionary(&slice)
	fmt.Println(*result)
}
