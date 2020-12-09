package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFilename := "input.txt"
	var inputs []int

	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineInt, _ := strconv.Atoi(line)
		inputs = append(inputs, lineInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var inputs2 []int
	copy(inputs2, inputs)

	for i := range inputs {
		for j := range inputs {
			if inputs[i]+inputs[j] == 2020 {
				fmt.Println("FOUND PAIR!: ", inputs[i], inputs[j])
				fmt.Println("Product of both: ", inputs[i]*inputs[j])
			}
		}
	}

	for i := range inputs {
		for j := range inputs {
			for k := range inputs {
				if inputs[i]+inputs[j]+inputs[k] == 2020 {
					fmt.Println("FOUND TRIPLE!: ", inputs[i], inputs[j], inputs[k])
					fmt.Println("Product of three: ", inputs[i]*inputs[j]*inputs[k])
				}
			}
		}
	}
}
