package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		items := make(map[string]int)
		input := scanner.Text()
		comp1 := strings.Split(input[:(len(input)/2)], "")
		comp2 := strings.Split(input[(len(input)/2):], "")

		for _, elem := range comp1 {
			r := []rune(elem)
			items[elem] = charToInt(r[0])
		}
		for _, elem := range comp2 {
			if val, exists := items[elem]; exists {
				delete(items, elem)
				// fmt.Println(elem)
				sum += val
				break
			}
		}

	}
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
