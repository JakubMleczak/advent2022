package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func overlap(worker1 []int, worker2 []int) bool {
	if worker1[1] >= worker2[0] && worker1[0] <= worker2[0] {
		return true
	}
	if worker1[1] >= worker2[1] && worker1[0] <= worker2[1] {
		return true
	}
	if worker1[1] <= worker2[1] && worker1[0] >= worker2[0] {
		return true
	}
	if worker2[1] <= worker1[1] && worker2[0] >= worker1[0] {
		return true
	}
	return false
}
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	nrWithin := 0
	nrWithin2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		workers := strings.Split(scanner.Text(), ",")

		worker1 := strings.Split(workers[0], "-")
		worker2 := strings.Split(workers[1], "-")

		ints1 := make([]int, len(worker1))
		for i, s := range worker1 {
			ints1[i], _ = strconv.Atoi(s)
		}

		ints2 := make([]int, len(worker2))
		for i, s := range worker2 {
			ints2[i], _ = strconv.Atoi(s)
		}

		if (ints1[0] >= ints2[0] && ints1[1] <= ints2[1]) || ints2[0] >= ints1[0] && ints2[1] <= ints1[1] {
			nrWithin++
		}
		if overlap(ints1, ints2) {
			nrWithin2++
		}
	}
	fmt.Println(nrWithin2)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
