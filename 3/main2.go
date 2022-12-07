package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func charToInt(c rune) int {
	if int(c) > 90 {
		// lower case
		return int(c) - 96
	} else {
		// upper case
		return int(c) - 38
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	for i := 0; i < len(input)-2; i += 3 {

		items1 := make(map[rune]int)
		items2 := make(map[rune]int)

		for _, char := range input[i] {
			items1[char] = charToInt(char)
		}
		for _, char := range input[i+1] {
			items2[char] = charToInt(char)
		}

		for _, elem := range input[i+2] {
			if _, exists := items1[elem]; exists {
				if val, exists2 := items2[elem]; exists2 {
					sum += val
					fmt.Printf("score %c \n", elem)
					break
				}

			}
		}
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
