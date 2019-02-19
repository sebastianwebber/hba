package hba

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ParseReader reads the buffer and parse a collection of rules
func ParseReader(buf io.Reader) (*[]Rule, error) {
	return parseReader(buf)
}

func parseReader(buf io.Reader) (*[]Rule, error) {

	if buf == nil {
		return nil, errors.New("could not parse a nil buffer")
	}

	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanLines)

	var out []Rule
	for scanner.Scan() {

		currentLine := strings.Trim(scanner.Text(), " ")

		if currentLine == "" || strings.HasPrefix(currentLine, "#") {
			continue
		}

		newRule, err := parseLine(currentLine)

		if err != nil {
			return nil, fmt.Errorf("could not parse line: %v", err)
		}

		out = append(out, *newRule)
	}

	return &out, nil
}
