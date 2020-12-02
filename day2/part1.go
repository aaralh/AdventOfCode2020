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
	fmt.Print("Day 2")

	f, err := os.Open("./day2/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	validPasswordCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Text()
		subStrings := strings.Split(line, " ")
		rule := strings.Split(subStrings[0], "-")
		char := strings.Trim(subStrings[1], ":")
		password := subStrings[2]

		charInPassword := strings.Count(password, char)

		minCharCount, _ := strconv.Atoi(rule[0])
		maxCharCount, _ := strconv.Atoi(rule[1])

		if (minCharCount <= charInPassword && charInPassword <= maxCharCount) {
			validPasswordCount++
		}

	}

	fmt.Println(validPasswordCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


}
