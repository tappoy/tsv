package tsv

import (
	"testing"
	"strings"
	"reflect"
	"os"
	)

const testData1 = `# comment test

L3F00	L3F01	L3F02
L4F10	L4F11	L4F12

L6F20		L6F22	L6F23
`

var wantTsvLines1 = []TsvLine{
	TsvLine{3, "L3F00	L3F01	L3F02", []string{"L3F00", "L3F01", "L3F02"}},
	TsvLine{4, "L4F10	L4F11	L4F12", []string{"L4F10", "L4F11", "L4F12"}},
	TsvLine{6, "L6F20		L6F22	L6F23", []string{"L6F20", "", "L6F22", "L6F23"}},
}

func TestParseTsv(t *testing.T) {
	r := strings.NewReader(testData1)
	lines, err := ParseTsv(r)
	if err != nil {
		t.Errorf("ParseTsv: %v", err)
	}
	if !reflect.DeepEqual(lines, wantTsvLines1) {
		t.Errorf("ParseTsv: got %v, want %v", lines, wantTsvLines1)
	}
}

type errorReader struct{}
func (errorReader) Read(p []byte) (n int, err error) {
	return 0, os.ErrInvalid
}

func TestParseTsvError(t *testing.T) {
	_, err := ParseTsv(errorReader{})
	if err == nil {
		t.Errorf("ParseTsv: got nil, want error")
	}
}

func TestReadTsvFile(t *testing.T) {
	// write wantTsvLines1 to a file
	f, err := os.Create("tmp/test.tsv")
	if err != nil {
		t.Errorf("Create: %v", err)
	}
	f.WriteString(testData1)
	f.Close()

	got, err := ReadTsvFile("tmp/test.tsv")
	if err != nil {
		t.Errorf("ReadTsv: %v", err)
	}

	if !reflect.DeepEqual(got, wantTsvLines1) {
		t.Errorf("ReadTsv: got %v, want %v", got, wantTsvLines1)
	}
}

func TestReadTsvFileFail(t *testing.T) {
	_, err := ReadTsvFile("nonexistent.tsv")
	if err == nil {
		t.Errorf("ReadTsv: got nil, want error")
	}
}
