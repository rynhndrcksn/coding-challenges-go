package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	countFileBytes := flag.Bool("c", false, "Count bytes in file")
	countFileLines := flag.Bool("l", false, "Count lines in file")
	countFileWords := flag.Bool("w", false, "Count words in file")
	countFileChars := flag.Bool("m", false, "Count chars in file")
	flag.Parse()

	filename := flag.Arg(0)

	// if no flags are passed, assume we want the byte, line, and word count
	if !*countFileBytes && !*countFileLines && !*countFileWords && !*countFileChars {
		*countFileBytes = true
		*countFileLines = true
		*countFileWords = true
	}

	if *countFileBytes {
		counter(filename, "bytes")
	}
	if *countFileLines {
		counter(filename, "lines")
	}
	if *countFileWords {
		counter(filename, "words")
	}
	if *countFileChars {
		counter(filename, "chars")
	}
}

// private helper func to count how many bytes, lines, words, or chars are in a file
func counter(filename string, s string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			msg := fmt.Sprintf("Problem closing file: %v", err)
			panic(msg)
		}
	}(file)

	count := 0
	scanner := bufio.NewScanner(file)
	if s == "bytes" {
		scanner.Split(bufio.ScanBytes)
	} else if s == "lines" {
		scanner.Split(bufio.ScanLines)
	} else if s == "words" {
		scanner.Split(bufio.ScanWords)
	} else if s == "chars" {
		scanner.Split(bufio.ScanRunes)
	}
	for scanner.Scan() {
		count++
	}
	fmt.Println(s, "=", count)
}
