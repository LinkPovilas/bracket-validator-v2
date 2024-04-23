package main

import (
	"log"
	"os"

	"github.com/LinkPovilas/bracket-validator-v2/validations"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide a file path as the first argument")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to open file:\n%s", err)
	}
	defer file.Close()

	validations.VerifyBracketPairsInFile(file, filePath)
}
