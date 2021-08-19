package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("Hyoon.txt")
	if err != nil {
		log.Println("FAILED TO READ FILE WITH PATH")
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // WE ASSUME EACH LINE IS LESS THAN 65536 CHARACTERS
	var hyoons []string
	for scanner.Scan() {
		hyoons = append(hyoons, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(hyoons, " && "))
}
