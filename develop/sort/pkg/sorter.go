package pkg

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type Sorter struct {
	textByLines  []string
	flags        []string
	targetColumn int
}

func splitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func InitSorter(t string, fs []string) (*Sorter, error) {
	if t == "" {
		return nil, errors.New("empty text")
	}

	lines := strings.Split(t, "\n")
	//lines := splitLines(t)
	return &Sorter{lines, fs, 0}, nil
}

func (s *Sorter) Start() error {
	for i := 0; i < len(s.flags); i++ {
		flag := s.flags[i]
		switch flag {
		case "-k":
			if i+1 == len(s.flags) {
				return errors.New("error of column number")
			}

			column, err := strconv.Atoi(s.flags[i+1])
			i++
			if err != nil {
				return err
			}
			s.targetColumn = column

			sortedText, err := sortByColumn(column, s.textByLines)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		case "-r":
			sortedText, err := reverseSort(s.textByLines)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		case "-u":
			sortedText, err := deleteOfDublicates(s.textByLines)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		case "-n":
			sortedText, err := sortByNumber(s.targetColumn, s.textByLines)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		case "-h":
			sortedText, err := sortByNumberWithSuffix(s.targetColumn, s.textByLines)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		case "-M":
			sortedText, err := sortByMonth(s.targetColumn, s.textByLines)
			if err != nil {
				return err
			}
			s.textByLines = sortedText

		case "-b":

		case "-c":
			return errors.New("-с is solo flag")

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

func (s *Sorter) CheckForSort() (bool, error) {
	sortedText, err := sortByColumn(s.targetColumn, s.textByLines)
	if err != nil {
		return false, err
	}
	return equalSlices(s.textByLines, sortedText), nil
}

func (s *Sorter) GetText() string {
	return strings.Join(s.textByLines, "\n")
}
