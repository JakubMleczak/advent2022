package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	containers := make([][]string, 9)

	containers[0] = append(containers[0], "F", "T", "C", "L", "R", "P", "G", "Q")
	containers[1] = append(containers[1], "N", "Q", "H", "W", "R", "F", "S", "J")
	containers[2] = append(containers[2], "F", "B", "H", "W", "P", "M", "Q")
	containers[3] = append(containers[3], "V", "S", "T", "D", "F")
	containers[4] = append(containers[4], "Q", "L", "D", "W", "V", "F", "Z")
	containers[5] = append(containers[5], "Z", "C", "L", "S")
	containers[6] = append(containers[6], "Z", "B", "M", "V", "D", "F")
	containers[7] = append(containers[7], "T", "J", "B")
	containers[8] = append(containers[8], "Q", "N", "B", "G", "L", "S", "P", "H")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//why this did not work Frankie ???
	r, _ := regexp.Compile("[1|2|3|4|5|6||7|8|9|0]+")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		match := r.FindAllString(input, -1)

		n, _ := strconv.Atoi(match[0])
		t, _ := strconv.Atoi(match[1])
		f, _ := strconv.Atoi(match[2])
		t--
		f--

		poppedVal := containers[t][len(containers[t])-n : len(containers[t])]
		// fmt.Println(poppedVal)
		containers[t] = containers[t][:len(containers[t])-n]
		containers[f] = append(containers[f], poppedVal...)

	}
	fmt.Println(containers)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
