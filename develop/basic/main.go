package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	options := ntp.QueryOptions{Timeout: 30 * time.Second, TTL: 10}
	response, err := ntp.QueryWithOptions("0.beevik-ntp.pool.ntp.org", options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = response.Validate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	time := time.Now().Add(response.ClockOffset).Format("15:04:05")
	fmt.Fprintln(os.Stderr, time)
	os.Exit(1)
}
