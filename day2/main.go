package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFilename := "input.txt"

	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
