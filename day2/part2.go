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

		firstCharPos, _ := strconv.Atoi(rule[0])
		secondCharPos, _ := strconv.Atoi(rule[1])

		firstCharIndex := firstCharPos - 1
		secondCharIndex := secondCharPos - 1

		if string(password[firstCharIndex]) == char && string(password[secondCharIndex]) != char {
			validPasswordCount++
		}

		if string(password[firstCharIndex]) != char && string(password[secondCharIndex]) == char {
			validPasswordCount++
		}

	}

	fmt.Println(validPasswordCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


}
