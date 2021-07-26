package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mysort/pkg"
)

func main() {
	//dir := os.Args[1]
	//fmt.Println(dir)

	r, _ := io.Pipe()
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	data := buf.String()
	fmt.Print(data)

	//file, err := os.Open("develop/sort/factory.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//data := make([]byte, 100)
	//_, err = file.Read(data)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	text := string(data)
	flags := flag.Args()

	sorter, err := pkg.InitSorter(text, flags)
	if err != nil {
		println(err)
	}

	if len(flags) == 0 {
		println("empty flags")
	} else if len(flags) == 1 && flag.Arg(1) == "-c" {
		println(sorter.CheckForSort())
	} else {
		err = sorter.Start()
		if err != nil {
			println(err)
		} else {
			newText := sorter.GetText()
			println(newText)
		}
	}
}
