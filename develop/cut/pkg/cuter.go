package pkg

import (
	"encoding/json"
	"errors"
)

type Cuter struct {
	textByLines                  []string
	flags                        []string
	resultLines                  []string
	outputLinesOnlyWithSeparator bool
	delimiter                    string
	targetFields                 []int
}

func InitCuter(lines []string, fs []string) (*Cuter, error) {
	if lines == nil {
		return nil, errors.New("empty text")
	}
	return &Cuter{lines, fs, nil, false, "\t", nil}, nil
}

func (c *Cuter) GetText() string {
	var s string
	for index, line := range c.resultLines {
		if index == 0 {
			s += line
		} else {
			s += "\n" + line
		}
	}
	return s
}

func (c *Cuter) Start() error {
	err := c.switchFlags()
	if err != nil {
		return err
	}

	if c.targetFields != nil {
		if c.delimiter == "" {
			c.resultLines = c.textByLines
		} else {
			needFields := cutFields(c.textByLines, c.delimiter, c.targetFields, c.outputLinesOnlyWithSeparator)
			c.resultLines = needFields
		}
	} else {
		return errors.New("input fields")
	}

	return nil
}

func (c *Cuter) switchFlags() error {
	for i := 0; i < len(c.flags); i++ {
		flag := c.flags[i]
		switch flag {
		case "-d":
			if i+1 < len(c.flags) {
				c.delimiter = c.flags[i+1]
				i++
			} else {
				return errors.New("input delimiter")
			}

		case "-s":
			c.outputLinesOnlyWithSeparator = true

		case "-f":
			if i+1 < len(c.flags) {
				fieldsSliceString := "[" + c.flags[i+1] + "]"
				var fieldsIntSlice []int
				if err := json.Unmarshal([]byte(fieldsSliceString), &fieldsIntSlice); err != nil {
					panic(err)
				}
				i++
				c.targetFields = fieldsIntSlice
			} else {
				return errors.New("input fields")
			}

		default:
			return errors.New("unknown flag: " + c.flags[i])
		}
	}

	return nil
}
