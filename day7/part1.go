package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	name string
	amount int
}

func calculateBagsContainingGoldBags(bagMap map[string][]bag) int {
	tmpMap := make(map[string]bool)

	for bagName, _ := range bagMap {
		if !tmpMap[bagName] {
			tmpMap[bagName] = false
		}
	}

	for i := 0; i < 10; i++ {
		for bagName, contents := range bagMap {
			for _, bag := range contents {
				if (bag.name == "shiny gold") {
					tmpMap[bagName] = true
					continue
				} else if tmpMap[bag.name] {
					tmpMap[bagName] = true
					continue
				}
			}
		}
	}
	goldCount := 0
	for _, val := range tmpMap {
		if (val) {
			goldCount++
		}
	}
	return goldCount
}

func main() {
	fmt.Println("Day 7")

	f, err := os.Open("./day7/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	bagMap := make(map[string][]bag)
	// plaid salmon bags contain 1 faded coral bag, 4 clear lavender bags, 5 wavy tan bags.
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "contain ")
		nameParts := strings.Split(strings.TrimSpace(parts[0]), " ")
		bagName := strings.Join(nameParts[:len(nameParts)-1], " ")
		bagContents := strings.Split(parts[1], ", ")
		var containsBags []bag
		if (bagContents[0] != "no other bags.") {
			for _, content := range bagContents {
				parts := strings.Split(content, " ")
				amount, _ := strconv.Atoi(parts[0])
				name := strings.TrimSpace(strings.Join(parts[1:3], " "))
				bag := bag{name:name, amount: amount}
				containsBags = append(containsBags, bag)
			}
		}
		bagMap[strings.TrimSpace(bagName)] = containsBags
	}
	count := calculateBagsContainingGoldBags(bagMap)
	fmt.Println("Result: ", count)
}