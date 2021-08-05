package pkg

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func Test_createFile(t *testing.T) {
	tables := []struct {
		fullUrlFile string
		fileName    string
	}{
		{"https://www.google.com/", "google"},
		{"https://www.google.com/", ""},
		{"https://www.google.com/qwe", ""},
		{"https://www.google.com/qwe", "nf"},
	}

	for _, table := range tables {
		filePath, err := createFile(table.fullUrlFile, table.fileName)

		assert.NoError(t, err)
		assert.Equal(t, filePath, filePath)
		_, err = os.Stat(filePath)
		assert.NoError(t, err)
		err = os.RemoveAll(strings.Split(filePath, "/")[0])
		assert.NoError(t, err)
	}
}
