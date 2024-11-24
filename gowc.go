package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	lineFlag := flag.Bool("l", false, "Enable -l flag")
	wordFlag := flag.Bool("w", false, "Enable -w flag")
	charFlag := flag.Bool("m", false, "Enable -m flag")
	byteFlag := flag.Bool("c", false, "Enable -c flag")

	flag.Parse()
	args := flag.Args()

	var file *os.File
	if len(args) == 0 {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	if !*lineFlag && !*wordFlag && !*charFlag && !*byteFlag {
		*charFlag = true
		*lineFlag = true
		*wordFlag = true
	}

	var result []string

	if *lineFlag {
		file.Seek(0, 0)
		result = append(result, countLines(file))
	}
	if *wordFlag {
		file.Seek(0, 0)
		result = append(result, countWords(file))
	}
	if *charFlag {
		file.Seek(0, 0)
		result = append(result, countChars(file))
	}
	if *byteFlag {
		file.Seek(0, 0)
		result = append(result, countBytes(file))
	}

	if len(result) > 0 {
		if len(args) > 0 {
			fmt.Println(strings.Join(result, " "), args[0])
		} else {
			fmt.Println(strings.Join(result, " "))
		}
	}
}

func countBytes(file *os.File) string {
	buffer := make([]byte, 1024)
	totalBytes := int64(0)

	for {
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return "Error"
		}
		totalBytes += int64(n)

		if err != nil {
			break
		}
	}
	return strconv.FormatInt(totalBytes, 10)
}

func countLines(file *os.File) string {
	buffer := make([]byte, 1024)
	count := 0
	sep := []byte{'\n'}

	for {
		c, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return "Error"
		}
		count += bytes.Count(buffer[:c], sep)

		if err != nil {
			break
		}
	}
	return strconv.FormatInt(int64(count), 10)

}

func countWords(file *os.File) string {
	buffer := make([]byte, 1024)
	count := 0

	for {
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return "Error"
		}
		count += len(strings.Fields(string(buffer[:n])))

		if err != nil {
			break
		}
	}
	return strconv.FormatInt(int64(count), 10)
}

func countChars(file *os.File) string {
	buffer := make([]byte, 1024)
	count := 0

	for {
		c, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return "Error"
		}
		count += utf8.RuneCount(buffer[:c])

		if err != nil {
			break
		}
	}
	return strconv.FormatInt(int64(count), 10)
}
