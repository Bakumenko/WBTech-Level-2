package main

import (
	"bufio"
	"fmt"
	"io"
	"mytelnet/pkg"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	if len(os.Args) < 3 {
		println("input more data")
		return
	}

	if len(os.Args) > 4 {
		println("extra data")
		return
	}

	domainOrIp := os.Args[1]
	port := os.Args[2]

	//flagTimeout := flag.String("timeout", "10s","set timeout")
	//flag.Parse()
	//println(*flagTimeout)
	timeout := 10 * time.Second
	if len(os.Args) == 4 {
		timeNumber := os.Args[3]
		number, err := strconv.Atoi(timeNumber)
		if err != nil {
			println("error timeout time")
		}
		timeout = time.Duration(number) * time.Second
	}

	telneter, err := pkg.InitTelneter(domainOrIp, port, timeout)
	if err != nil {
		println(err.Error())
		return
	}

	stat, err := os.Stdin.Stat()
	if (stat.Mode() & os.ModeNamedPipe) == 0 {
		println("input data for server")
		return
	}

	var lines []string
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
		if err != nil {
			println(err.Error())
			return
		}
	}

	response, err := telneter.Send(lines)
	if err != nil {
		println(err.Error())
	}

	_, err = fmt.Fprintln(os.Stdout, response)
	if err != nil {
		println(err)
	}
}
