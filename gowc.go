package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

	counts := ""

	if *byteFlag {
		counts += countBytes(file)
	}
	if *lineFlag {
		counts += countLines(file)
	}
	if *wordFlag {
		counts += countWords(file)
	}
	if *charFlag {
		counts += countChars(file)
	}
	fmt.Println(counts, filename)
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
	sep := []byte{' '}

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
