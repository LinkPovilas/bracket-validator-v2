package utils

import (
	"fmt"
	"strings"
)

func GetInvalidBracketMessage(char byte, filePath string, lineNumber, columnNumber int) string {
	fileNameIndex := len(strings.Split(filePath, "/")) - 1
	fileName := strings.Split(filePath, "/")[fileNameIndex]

	return fmt.Sprintf("Invalid bracket %c found at %s:%d:%d.\n", char, fileName, lineNumber, columnNumber)
}
