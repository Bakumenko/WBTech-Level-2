package main

import (
	"github.com/mitchellh/go-ps"
)

func main() {
	processList, err := ps.Processes()
	if err != nil {
		println("ps.Processes() Failed, are you using windows?")
		return
	}

	println("PID  |  COMMAND")
	println("---------------")
	for x := range processList {
		var process ps.Process
		process = processList[x]
		print(process.Pid())
		print("    |  ")
		println(process.Executable())
	}
}
