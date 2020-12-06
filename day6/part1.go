package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	fmt.Println("Day 6")

	f, err := os.Open("./day6/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var groupAnswers []int

	var currentGroupAnswers []string
	for scanner.Scan() {
		line := scanner.Text()
		if (line == "") {
			fmt.Println()
			currentGroupAnswers = unique(currentGroupAnswers)
			groupAnswers = append(groupAnswers, len(currentGroupAnswers))
			currentGroupAnswers = nil
			continue
		}
		answers := strings.Split(line, "")
		currentGroupAnswers = append(currentGroupAnswers, answers...)
	}

	currentGroupAnswers = unique(currentGroupAnswers)
	groupAnswers = append(groupAnswers, len(currentGroupAnswers))
	currentGroupAnswers = nil

	fmt.Println(groupAnswers)
	result := 0

	for _, item := range groupAnswers {
		result += item
	}

	fmt.Println("Result:", result)
}
