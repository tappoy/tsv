// This package is used to read TSV(Tab Separated Values) files.
// If the first character of a line is '#', it becomes a comment.
// Blank lines are ignored.
package tsv

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type TsvLine struct {
	LineNo int
	Line   string
	Fields []string
}

// ParseTsv reads a TSV stream and returns a slice of TsvLine.
//
// Errors:
//   - bufio.Scanner.Err()
func ParseTsv(r io.Reader) ([]TsvLine, error) {
	var lines []TsvLine
	scanner := bufio.NewScanner(r)
	lineNo := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNo++
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		fields := strings.Split(line, "\t")
		lines = append(lines, TsvLine{lineNo, line, fields})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// ReadTsvFile reads a TSV file and returns a slice of TsvLine.
//
// Errors:
//   - os.Open()
//   - bufio.Scanner.Err()
func ReadTsvFile(filename string) ([]TsvLine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ParseTsv(file)
}
