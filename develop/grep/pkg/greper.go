package pkg

import (
	"bufio"
	"errors"
	"strings"
)

type Greper struct {
	textByLines  []string
	flags        []string
	targerString string
	ignoreCase   bool
	lineNumber   bool
}

func splitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func InitGreper(t string, fs []string, targerString string) (*Greper, error) {
	if t == "" {
		return nil, errors.New("empty text")
	}

	lines := strings.Split(t, "\n")
	//lines := splitLines(t)
	return &Greper{lines, fs, targerString, false, false}, nil
}

func (g *Greper) Start() error {
	for i := 0; i < len(g.flags); i++ {
		flag := g.flags[i]
		switch flag {

		default:
			return errors.New("unknown flag")
		}
	}
	return nil
}

func (g *Greper) GetText() string {
	return strings.Join(g.textByLines, "\n")
}
