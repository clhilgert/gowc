package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	bytesFlag := flag.Bool("c", false, "Enable -c flag")
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

	if *bytesFlag {
		countBytes(file, filename)
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
