package main

import (
	"flag"
	"io/ioutil"
	"net"
	"os"
	"sync"
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
	address := net.JoinHostPort(flagArgs[0], flagArgs[1])

	c := NewTelnetClient(address, timeout, ioutil.NopCloser(os.Stdin), os.Stdout)
	println("Connected to " + address)
	if err := c.Connect(); err != nil {
		println(err.Error())
	}
	defer func() {
		err := c.Close()
		println(err.Error())
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := c.Send(); err != nil {
				println(err.Error())
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := c.Receive(); err != nil {
				println(err.Error())
				return
			}
		}
	}()

	wg.Wait()
}
