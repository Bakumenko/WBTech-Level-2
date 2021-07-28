package main

import (
	"bufio"
	"io"
	"mygrep/pkg"
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

func main() {
	if len(os.Args) < 3 {
		println("input more data")
		return
	}
	targetString := os.Args[1]
	flags := os.Args[2:]

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
	}

	lines, err := readLinesFromSource(sourceForRead)
	if err != nil {
		println(err.Error())
		return
	}

	sorter, err := pkg.InitGreper(lines, flags, targetString)
	if err != nil {
		println(err.Error())
		return
	}

	err = sorter.Start()
	if err != nil {
		println(err.Error())
		return
	}

	result := sorter.GetText()
	println(result)
}
