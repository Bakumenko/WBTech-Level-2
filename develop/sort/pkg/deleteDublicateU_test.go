package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteOfDublicates(t *testing.T) {
	tables := []struct {
		lines []string
		res   []string
	}{
		{[]string{"asd", "asd"}, []string{"asd"}},
		{[]string{"asd"}, []string{"asd"}},
		{[]string{"", ""}, []string{""}},
		{[]string{"asd", "qwe"}, []string{"asd", "qwe"}},
		{[]string{"asd", "asd", "qwe"}, []string{"asd", "qwe"}},
		{nil, []string{}},
	}
	for _, table := range tables {
		result := deleteOfDublicates(table.lines)

		assert.Equal(t, result, table.res)
	}
}
