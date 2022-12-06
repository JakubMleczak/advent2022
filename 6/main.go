package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(line)-14-1; i++ {
		sameChar := false
		key := line[i : i+14]
		// fmt.Print(key + " ")
		for j := 0; j < 14-1; j++ {

			for k := 1; k < 14-j; k++ {

				if key[j] == key[j+k] {
					sameChar = true
				}
			}
		}
		if !sameChar {
			fmt.Println(i)
			break
		}
	}
}
