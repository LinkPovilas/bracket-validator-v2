package utils

import (
	"fmt"
	"testing"
)

type testpair struct {
	Name         string
	Char         byte
	FilePath     string
	LineNumber   int
	ColumnNumber int
}

var tests = [2]testpair{
	{
		Name:         "ShortFilePath",
		Char:         '(',
		FilePath:     "./test-file.txt",
		LineNumber:   10,
		ColumnNumber: 6,
	},
	{
		Name:         "LongFilePath",
		Char:         ']',
		FilePath:     "./src/test-data/test-file.txt",
		LineNumber:   21,
		ColumnNumber: 7,
	},
}

func TestGetInvalidBracket(t *testing.T) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			want := fmt.Sprintf(
				"Invalid bracket %c found at test-file.txt:%d:%d.\n",
				test.Char,
				test.LineNumber,
				test.ColumnNumber,
			)
			got := GetInvalidBracketMessage(
				test.Char,
				test.FilePath,
				test.LineNumber,
				test.ColumnNumber,
			)

			if got != want {
				t.Errorf("Expected output: %s, but got: %s", want, got)
			}
		})
	}

}
