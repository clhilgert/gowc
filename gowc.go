package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func main() {

	byteFlag := flag.Bool("c", false, "Enable -c flag")
	lineFlag := flag.Bool("l", false, "Enable -l flag")
	wordFlag := flag.Bool("w", false, "Enable -w flag")
	charFlag := flag.Bool("m", false, "Enable -w flag")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("No filename provided")
		return
	}
	filename := args[0]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if *byteFlag {
		countBytes(file, filename)
	}
	if *lineFlag {
		countLines(file, filename)
	}
	if *wordFlag {
		countWords(file, filename)
	}
	if *charFlag {
		countChars(file, filename)
	}
}

func countBytes(file *os.File, filename string) {
	buffer := make([]byte, 1024)
	totalBytes := int64(0)

	for {
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return
		}
		totalBytes += int64(n)

		if err != nil {
			break
		}
	}
	fmt.Println(totalBytes, filename)
}

func countLines(file *os.File, filename string) {
	buffer := make([]byte, 1024)
	count := 0
	sep := []byte{'\n'}

	for {
		c, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return
		}
		count += bytes.Count(buffer[:c], sep)

		if err != nil {
			break
		}
	}
	fmt.Println(count, filename)
}

func countWords(file *os.File, filename string) {
	buffer := make([]byte, 1024)
	count := 0
	sep := []byte{' '}

	for {
		c, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return
		}
		count += bytes.Count(buffer[:c], sep)

		if err != nil {
			break
		}
	}
	fmt.Println(count, filename)
}

func countChars(file *os.File, filename string) {
	buffer := make([]byte, 1024)
	count := 0

	for {
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return
		}
		count += utf8.RuneCount(buffer[:n])

		if err != nil {
			break
		}
	}
	fmt.Println(count, filename)
}
