package pkg

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Wgeter struct {
	url      string
	fileName string
}

func InitWgeter(url string, filename string) (*Wgeter, error) {
	if url == "" {
		return nil, errors.New("empty text")
	}
	return &Wgeter{url, filename}, nil
}

func (w *Wgeter) Start() error {
	fullURLFile := w.url

	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		return err
	}
	path := fileURL.Path
	if w.fileName == "" {
		segments := strings.Split(path, "/")
		w.fileName = segments[len(segments)-1]
	}

	file, err := os.Create(w.fileName)
	if err != nil {
		return err
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fullURLFile)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	println("Downloaded a file " + w.fileName + " with size " + strconv.FormatInt(size, 10))
	return nil
}
