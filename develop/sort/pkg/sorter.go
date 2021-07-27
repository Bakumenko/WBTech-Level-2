package pkg

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Sorter struct {
	textByLines     []string
	flags           []string
	targetColumns   []int
	neededToReverse bool
	checkForSort    string
}

func InitSorter(lines []string, fs []string) (*Sorter, error) {
	if lines == nil {
		return nil, errors.New("empty text")
	}
	return &Sorter{lines, fs, nil, false, ""}, nil
}

func (s *Sorter) GetText() string {
	if s.checkForSort != "" {
		return s.checkForSort
	} else {
		return strings.Join(s.textByLines, "\n")
	}
}

func (s *Sorter) Start() error {
	var flagsWithoutPreSortedFlags []string
	for i := 0; i < len(s.flags); i++ {
		flag := s.flags[i]
		switch flag {
		case "-r":
			s.neededToReverse = true

		case "-u":
			linesWithoutDublicate, err := deleteOfDublicates(s.textByLines)
			if err != nil {
				return err
			}
			s.textByLines = linesWithoutDublicate

		case "-c":
			if len(s.flags) != 1 {
				return errors.New("-c is solo flag")
			}

			if equalSlices(s.textByLines, simpleSort(s.textByLines, false)) {
				s.checkForSort = "sorted"
			} else {
				s.checkForSort = "not sorted"
			}

		case "-k":
			columns := s.flags[i+1]
			i++
			sliceColumns, err := parseNumberOfColumnsFromString(columns)
			if err != nil {
				return err
			}
			s.targetColumns = sliceColumns

		default:
			flagsWithoutPreSortedFlags = append(flagsWithoutPreSortedFlags, flag)
		}
	}

	s.flags = flagsWithoutPreSortedFlags
	//if len(s.flags) == 0 {
	//	sort.Strings(s.textByLines)
	//	return nil
	//}

	if len(s.flags) > 2 {
		return errors.New("incompatible flags" + strings.Join(s.flags, " "))
	} else if len(s.flags) == 0 {
		if s.targetColumns != nil {
			sortedLines, err := sortByColumn(s.targetColumns, s.textByLines, s.neededToReverse)
			if err != nil {
				return err
			}
			s.textByLines = sortedLines
		} else {
			s.textByLines = simpleSort(s.textByLines, s.neededToReverse)
		}
	} else {
		lastFlag := s.flags[0]
		switch lastFlag {
		case "-n":
			if len(s.targetColumns) == 1 {
				sortedText, err := sortByNumber(s.targetColumns[0], s.textByLines, s.neededToReverse)
				if err != nil {
					return err
				}
				s.textByLines = sortedText
			}

		case "-h":
			sortedText, err := sortByNumberWithSuffix(s.targetColumns[0], s.textByLines, s.neededToReverse)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		case "-M":
			sortedText, err := sortByMonth(s.targetColumns[0], s.textByLines, s.neededToReverse)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		default:
			return errors.New("unknown flag")
		}
	}
	return nil
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func simpleSort(lines []string, neededToReverse bool) []string {
	if neededToReverse {
		sort.Sort(sort.Reverse(sort.StringSlice(lines)))
	} else {
		sort.Strings(lines)
	}
	return lines
}

func parseNumberOfColumnsFromString(s string) ([]int, error) {
	var columns []int
	if strings.Contains(s, ",") {
		sliceStringColumns := strings.Split(s, ",")
		if len(sliceStringColumns) > 2 {
			return nil, errors.New("incorrect columns")
		}

		firstColumn, err := strconv.Atoi(sliceStringColumns[0])
		if err != nil {
			return nil, err
		}

		lastColumn, err := strconv.Atoi(sliceStringColumns[1])
		if err != nil {
			return nil, err
		}
		for i := firstColumn - 1; i < lastColumn-1; i++ {
			columns = append(columns, i)
		}
	} else {
		column, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	sort.Ints(columns)
	return columns, nil
}
