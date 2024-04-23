package validations

import (
	"bufio"
	"fmt"
	"os"

	"github.com/LinkPovilas/bracket-validator-v2/utils"
)

type Bracket struct {
	Char     byte
	Line     int
	Position int
}

func VerifyBracketPairsInFile(file *os.File, filePath string) {
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	brackets := []Bracket{}

	for scanner.Scan() {
		line := scanner.Bytes()

		for i, char := range line {
			switch char {
			case '(', '[', '{':
				brackets = append(brackets, Bracket{char, lineNumber, i + 1})
			case ')', ']', '}':
				if len(brackets) == 0 {
					fmt.Println(utils.GetInvalidBracketMessage(char, filePath, lineNumber, i+1))
				} else {
					last := len(brackets) - 1
					if (char == ')' && brackets[last].Char == '(') ||
						(char == ']' && brackets[last].Char == '[') ||
						(char == '}' && brackets[last].Char == '{') {
						brackets = brackets[:last]
					} else {
						fmt.Println(utils.GetInvalidBracketMessage(char, filePath, lineNumber, i+1))
					}
				}
			}
		}
		lineNumber++
	}

	for _, b := range brackets {
		fmt.Println(utils.GetInvalidBracketMessage(b.Char, filePath, b.Line, b.Position))
	}
}
