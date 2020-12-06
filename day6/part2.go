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
	persons := 0
	var currentGroupAnswers []string
	for scanner.Scan() {
		line := scanner.Text()
		if (line == "") {
			fmt.Println()
			answers := unique(currentGroupAnswers)
			groupAnswerString := strings.Join(currentGroupAnswers[:], "")
			answerCount := 0
			for _, answer := range answers {
				count := strings.Count(groupAnswerString, answer)
				if (count == persons) {
					answerCount++
				}
			}

			groupAnswers = append(groupAnswers, answerCount)
			currentGroupAnswers = nil
			persons = 0
			continue
		}
		persons += 1
		answers := strings.Split(line, "")
		currentGroupAnswers = append(currentGroupAnswers, answers...)
	}

	answers := unique(currentGroupAnswers)
	groupAnswerString := strings.Join(currentGroupAnswers[:], "")
	answerCount := 0
	for _, answer := range answers {
		count := strings.Count(groupAnswerString, answer)
		if (count == persons) {
			answerCount++
		}
	}

	groupAnswers = append(groupAnswers, answerCount)

	fmt.Println(groupAnswers)
	result := 0

	for _, item := range groupAnswers {
		result += item
	}

	fmt.Println("Result:", result)
}
