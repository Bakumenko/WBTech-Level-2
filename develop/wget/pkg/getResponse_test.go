package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getResponse(t *testing.T) {
	urls := []string{
		"https://golang.org/",
		"https://www.google.ru/",
		"https://www.google.com/",
		"https://yandex.ru/",
		"https://yandex.com/",
	}

	for _, url := range urls {
		res, err := getResponse(url)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}
