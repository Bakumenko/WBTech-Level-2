package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_grepLinesNearSearched(t *testing.T) {
	tables := []struct {
		textByLines  []string
		targerString string
		before       int
		after        int
		ignoreCase   bool
		lineNumber   bool
		regular      bool
		invert       bool
		resLines     []string
		resCount     int
	}{
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, false, false, false, false, []string{"asd"}, 0},

		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, true, false, false, false, []string{"asd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, false, true, false, false, []string{"asd", "aSd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, false, false, true, false, []string{"1. asd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, false, false, false, true, []string{"1. asd"}, 0},

		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, true, true, false, false, []string{"asd", "aSd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, true, false, true, false, []string{"1. asd", "2. aSd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, true, false, false, true, []string{"1. asd", "2. aSdd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, false, true, true, false, []string{"asd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, false, true, false, true, []string{"asd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, false, false, true, true, []string{"asd"}, 0},

		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, true, true, true, false, []string{"asd", "aSd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, true, true, false, true, []string{"1. asd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, true, false, true, true, []string{"1. asd"}, 0},
		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, false, true, true, true, []string{"asd", "aSd"}, 0},

		{[]string{"asd", "aSd", "qwe"}, "asd", 0, 0, true, true, true, true, []string{"1. asd", "2. aSd"}, 0},
	}
	for _, table := range tables {
		resultLines, resultCount, _ := grepLinesNearSearched(table.textByLines, table.targerString, table.before, table.after, table.ignoreCase, table.lineNumber, table.regular, table.invert)

		assert.Equal(t, resultLines, table.resLines)
		assert.Equal(t, resultCount, table.resCount)
	}
}
