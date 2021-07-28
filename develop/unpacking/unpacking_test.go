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
		{`q\413`, "q4444444444444", nil},
		{"a0b2c12", "bbcccccccccccc", nil},
		{`qw\\\e`, "", errors.New("invalid string")},
	}
	for _, table := range tables {
		totalS, totalErr := UnpackingString(table.s)

		if totalS != table.res || totalErr == nil && table.err != nil || totalErr != nil && table.err == nil {
			t.Errorf("UnpackingString of (%v) was incorrect, got: (%v), want: (%v) with error: (%v), wanted error: (%v)\n",
				table.s, totalS, table.res, totalErr, table.err)
		}
	}
}
