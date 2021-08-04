package pkg

import (
	"fmt"
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
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, false, false, false, []string{"oleg", "ert", "rty", "QWE"}, 2},

		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, false, false, false, []string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, 3},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, true, false, false, []string{"1. oleg", "2. ert", "3. rty", "4. QWE"}, 2},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, false, true, false, []string{"oleg", "ert", "rty", "QWE"}, 2},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, false, false, true, []string{"ert", "rty", "QWE", "asd", "zxc"}, 4},

		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, true, false, false, []string{"1. oleg", "2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 3},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, false, true, false, []string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, 3},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, false, false, true, []string{"ert", "rty", "QWE", "asd", "zxc"}, 3},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, true, true, false, []string{"1. oleg", "2. ert", "3. rty", "4. QWE"}, 2},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, true, false, true, []string{"2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 4},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, false, true, true, []string{"ert", "rty", "QWE", "asd", "zxc"}, 4},

		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, true, true, false, []string{"1. oleg", "2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 3},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, true, false, true, []string{"2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 3},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, false, true, true, []string{"ert", "rty", "QWE", "asd", "zxc"}, 3},
		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, false, true, true, true, []string{"2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 4},

		{[]string{"oleg", "ert", "rty", "QWE", "asd", "zxc"}, "e", 1, 2, true, true, true, true, []string{"2. ert", "3. rty", "4. QWE", "5. asd", "6. zxc"}, 3},
	}
	for _, table := range tables {
		resultLines, resultCount, _ := grepLinesNearSearched(table.textByLines, table.targerString, table.before, table.after, table.ignoreCase, table.lineNumber, table.regular, table.invert)

		inputData := fmt.Sprintf("lines = (%v)\ntarget = (%v)\ncount before = %v, count after = %v\nignore case = %v, print line = %v, use regular = %v, invert = %v",
			table.textByLines, table.targerString, table.before, table.after, table.ignoreCase, table.lineNumber, table.regular, table.invert)
		assert.Equal(t, resultLines, table.resLines, inputData)
		assert.Equal(t, resultCount, table.resCount, inputData)
	}
}
