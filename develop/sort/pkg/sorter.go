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
	targetColumn    int
	neededToReverse bool
	checkForSort    string
}

func InitSorter(lines []string, fs []string) (*Sorter, error) {
	if lines == nil {
		return nil, errors.New("empty text")
	}
	return &Sorter{lines, fs, -1, false, ""}, nil
}

func (s *Sorter) GetText() string {
	if s.checkForSort != "" {
		return s.checkForSort
	} else {
		return strings.Join(s.textByLines, "\n")
	}
}

func (s *Sorter) Start() error {
	err := s.switchFlags()
	if err != nil {
		return err
	}

	if len(s.flags) >= 2 {
		return errors.New("incompatible flags: " + strings.Join(s.flags, " "))
	} else if len(s.flags) == 0 {
		if s.targetColumn != -1 {
			sortedLines, err := sortByColumn(s.targetColumn, s.textByLines, s.neededToReverse)
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
			sortedText, err := sortByNumber(s.targetColumn, s.textByLines, s.neededToReverse)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		case "-h":
			sortedText, err := sortByNumberWithSuffix(s.targetColumn, s.textByLines, s.neededToReverse)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		case "-M":
			sortedText, err := sortByMonth(s.targetColumn, s.textByLines, s.neededToReverse)
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
	var lowerLinesSlice []string
	for _, line := range lines {
		lowerLinesSlice = append(lowerLinesSlice, line)
	}
	if neededToReverse {
		sort.Sort(sort.Reverse(sort.StringSlice(lowerLinesSlice)))
	} else {
		sort.Strings(lowerLinesSlice)
	}
	return lowerLinesSlice
}

func (s *Sorter) switchFlags() error {
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
			if i+1 < len(s.flags) {
				column, err := strconv.Atoi(s.flags[i+1])
				if err != nil {
					return err
				}
				s.targetColumn = column - 1
				i++
			} else {
				return errors.New("input number of column")
			}

		default:
			flagsWithoutPreSortedFlags = append(flagsWithoutPreSortedFlags, flag)
		}
	}

	s.flags = flagsWithoutPreSortedFlags
	return nil
}
