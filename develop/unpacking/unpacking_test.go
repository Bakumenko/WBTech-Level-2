package main

import (
	"errors"
	"testing"
)

func TestUnpackingString(t *testing.T) {
	tables := []struct {
		s   string
		res string
		err error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("invalid string")},
		{"", "", nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
	}
	for _, table := range tables {
		totalS, totalErr := UnpackingString(table.s)
		if totalS != table.res || totalErr != table.err {
			t.Errorf("UnpackingString of (%v) was incorrect, got: %v, want: %v with error: %v, wanted error: %v\n",
				table.s, table.res, totalS, table.err, totalErr)
		}
	}
}
