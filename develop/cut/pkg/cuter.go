package pkg

import (
	"bufio"
	"encoding/json"
	"errors"
	"strings"
)

type Cuter struct {
	textByLines                  []string
	flags                        []string
	outputLinesOnlyWithSeparator bool
	delimiter                    string `defalult:"\t"`
	targetFields                 []int
	result                       []string
}

func splitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func InitCuter(t string, fs []string, targerString string) (*Cuter, error) {
	if t == "" {
		return nil, errors.New("empty text")
	}

	lines := strings.Split(t, "\n")
	//lines := splitLines(t)
	return &Cuter{textByLines: lines, flags: fs}, nil
}

func parseStringToSliceInt(s string) ([]int, error) {
	var is []int
	if err := json.Unmarshal([]byte(s), &is); err != nil {
		return nil, errors.New("error in fields")
	}
	return is, nil
}

func (c *Cuter) Start() error {
	for i := 0; i < len(c.flags); i++ {
		flag := c.flags[i]
		switch flag {
		case "-d":
			c.delimiter = c.flags[i+1]
			i++

		case "-s":
			c.outputLinesOnlyWithSeparator = true

		case "-f":
			fields, err := parseStringToSliceInt(c.flags[i+1])
			if err != nil {
				return err
			}
			i++
			result := chooseFields(c.textByLines, c.delimiter, fields, c.outputLinesOnlyWithSeparator)
			c.result = result
		default:
			return errors.New("unknown flag")
		}
	}
	return nil
}

func (c *Cuter) GetText() string {
	return strings.Join(c.textByLines, "\n")
}
