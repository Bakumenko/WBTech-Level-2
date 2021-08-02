package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getFieldForSortFromLines(t *testing.T) {
	tables := []struct {
		numberColumn int
		line         string
		res          string
	}{
		{0, "a b c d", "a"},
		{1, "a b c d", "b"},
		{2, "a b c d", "c"},
		{3, "a b c d", "d"},
		{4, "a b c d", "a"},
		{5, "a b c d", "a"},
		{3, "", ""},
		{-1, "", ""},
		{-1, "a b c d", "a b c d"},
		{2, "a   b    c      d", "c"},
	}
	for _, table := range tables {
		result := getFieldForSortFromLines(table.numberColumn, table.line)

		assert.Equal(t, result, table.res)
	}
}
