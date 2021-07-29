package main

import (
	"os"
	"os/user"
)

func main() {
	//for {
	//	line := waitForInputLine()
	//	switch {
	//	case strings.HasPrefix(line, "cd"):
	//		os.Chdir(strings.Split(line, " ")[1])
	//
	//	// ..check other builtins and special cases./
	//
	//	default:
	//		runBinary(line)
	//	}
	//}

	//curUser, err := user.Current()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//os.Chdir(curUser.HomeDir)
	//// or like this
	// cmd := exec.Command("cd", curUser.HomeDir)
	if len(os.Args) == 1 {
		curUser, err := user.Current()
		if err != nil {
			println(err.Error())
		}
		err = os.Chdir(curUser.HomeDir)
		if err != nil {
			println(err.Error())
		}
	} else {
		inputDirectory := os.Args[1]
		err := os.Chdir(inputDirectory)
		if err != nil {
			println(err.Error())
		}
	}

	path, err := os.Getwd()
	if err != nil {
		println(err.Error())
	}
	println(path)
}
