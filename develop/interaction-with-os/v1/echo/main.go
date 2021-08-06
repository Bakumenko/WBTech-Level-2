package main

import (
	"fmt"
	"os"
)

func main() {
	var inputText string
	if len(os.Args) == 1 {
		inputText = ""
	} else {
		inputText = os.Args[1]

	}
	_, err := fmt.Fprintf(os.Stdin, inputText)
	if err != nil {
		println(err.Error())
	}
	println()
}
