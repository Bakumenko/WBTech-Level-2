package main

import (
	"flag"
	"mytelnet/pkg"
	"time"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", time.Second*10, "connection timeout")
	flag.Parse()

	flagArgs := flag.Args()
	if len(flagArgs) != 2 {
		println("Missing required params, usage: \"go-telnet host port\" or \"go-telnet --timeout=10s host port\"")
	}

	host := flagArgs[0]
	port := flagArgs[1]

	telnet := pkg.NewTelneter(host, port, timeout)

	err := telnet.Start()
	if err != nil {
		println(err.Error())
	}

}
