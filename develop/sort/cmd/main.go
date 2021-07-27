package main

import (
	"bufio"
	"io"
	"mysort/pkg"
	"os"
	"strings"
)

func readLinesFromSource(in io.Reader) ([]string, error) {
	var lines []string
	reader := bufio.NewReader(in)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return lines, nil
}

func printLines(lines []string) {
	for _, line := range lines {
		println(line)
	}
}

func main() {
	flags := os.Args[1:]
	stat, err := os.Stdin.Stat()
	var sourceForRead io.Reader
	if err != nil {
		println(err.Error())
		return
	}

	if (stat.Mode() & os.ModeNamedPipe) != 0 {
		sourceForRead = os.Stdin
		println("read from stdin")
	} else {
		if len(flags) == 0 {
			println("input file name")
			return
		}
		fileName := flags[0]
		flags = flags[1:]
		file, err := os.Open(fileName)
		if err != nil {
			println(err.Error())
			return
		}
		sourceForRead = file
		println("read from file")
	}

	lines, err := readLinesFromSource(sourceForRead)
	if err != nil {
		println(err.Error())
		return
	}
	printLines(lines)
	printLines(flags)

	sorter, err := pkg.InitSorter(lines, flags)
	if err != nil {
		println(err)
		return
	}

	err = sorter.Start()
	if err != nil {
		println(err)
		return
	}

	result := sorter.GetText()
	println("result:" + result)
}
