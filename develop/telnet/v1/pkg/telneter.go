package pkg

import (
	"errors"
	"fmt"
	"github.com/reiver/go-telnet"
	"time"
)

type Telneter struct {
	url        string
	timeout    time.Duration
	connection *telnet.Conn
}

func InitTelneter(domainOrIp, port string, timeout time.Duration) (*Telneter, error) {
	url := domainOrIp + ":" + port
	if url == "" {
		return nil, errors.New("empty text")
	}

	println("Trying " + domainOrIp)
	conn, err := telnet.DialTo(url)
	if err != nil {
		return nil, err
	}
	println("Connected to " + url)

	return &Telneter{url, timeout, conn}, nil
}

func (t *Telneter) Send(lines []string) (string, error) {
	for _, line := range lines {
		_, err := t.connection.Write([]byte(line + "\n"))
		if err != nil {
			return "", err
		}
	}

	b := make([]byte, 100)
	_, err := t.connection.Read(b)
	if err != nil {
		return "", err
	}
	fmt.Println(string(b))
	return string(b), nil
}
