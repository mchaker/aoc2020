package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFilename := "input.txt"
	var inputLines []string

	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputLines = append(inputLines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	validPasswords := 0
	validPasswords2 := 0

	for _, line := range inputLines {
		if isValidPassword(line) == true {
			validPasswords++
		}

		if isValidPassword2(line) == true {
			validPasswords2++
		}
	}

	fmt.Println("valid passwords, part 1: ", validPasswords)
	fmt.Println("valid passwords, part 2: ", validPasswords2)
}

func isValidPassword(passLine string) bool {
	var validPassword bool

	splitStr := strings.Split(passLine, ":")
	policyStr := strings.Split(splitStr[0], " ")
	password := strings.TrimSpace(splitStr[1])
	letterStr := strings.TrimSpace(policyStr[1])
	repeatBoundsStr := strings.Split(policyStr[0], "-")
	minCount, _ := strconv.Atoi(repeatBoundsStr[0])
	maxCount, _ := strconv.Atoi(repeatBoundsStr[1])

	repeatCount := strings.Count(password, letterStr)

	if repeatCount >= minCount && repeatCount <= maxCount {
		validPassword = true
	} else {
		validPassword = false
	}

	return validPassword
}

func isValidPassword2(passLine string) bool {
	var validPassword bool

	splitStr := strings.Split(passLine, ":")
	policyStr := strings.Split(splitStr[0], " ")
	password := strings.TrimSpace(splitStr[1])
	letterStr := strings.TrimSpace(policyStr[1])
	validIndices := strings.Split(policyStr[0], "-")
	index1, _ := strconv.Atoi(validIndices[0])
	index2, _ := strconv.Atoi(validIndices[1])

	if (string(password[index1-1]) == letterStr) != (string(password[index2-1]) == letterStr) {
		validPassword = true
	} else {
		validPassword = false
	}

	return validPassword
}
