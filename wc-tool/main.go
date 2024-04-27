package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Define the flogs
	countBytes := flag.Bool("b", false, "Count bytes")
	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
  countChars := flag.Bool("c", false, "Count characters")	

	// Parse command-line flags
	flag.Parse()
  // Check if the filename is provided
	args := flag.Args()
	if len(args) == 0 {
		// Read from standard input
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading ffrom standard input: ", err)
			return
		}
		processData(data, *countBytes, *countLines, *countWords, *countChars)
	} else {
		// Read from file
		fileName := args[0]
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading file: ", err)
		}
		processData(data, *countBytes, *countLines, *countWords, *countChars)
	}

}

func processData(data []byte, countBytes, countLines, countWords, countChars bool) {

	byteCount := len(data)
	lineCount := strings.Count(string(data), "\n")
	wordCount := len(strings.Fields(string(data)))
	charCount := len([]rune(string(data)))

	if countBytes {
		fmt.Printf("byteCount: %8d", byteCount)
	}
	if countLines {
		fmt.Printf("lineCount: %8d", lineCount)
	}
	if countWords {
		fmt.Printf("wordCount: %8d", wordCount)
	}
	if countChars {
		fmt.Printf("charCount: %8d", charCount)
	}

	if countBytes || countLines || countWords || countChars {
		fmt.Println()
	} else {
		fmt.Printf("byteCount: %8d\nlineCount: %8d\nwordCount: %8d\ncharCount: %8d\n", byteCount, lineCount, wordCount, charCount)
	}
}