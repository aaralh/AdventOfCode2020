package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {
	fmt.Println("Day 3")

	f, err := os.Open("./day3/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	treeCount := 0
	lineCount := 0
	scanner := bufio.NewScanner(f)
	tree := "#"
	yPosition := 0

	for scanner.Scan() {
		line := scanner.Text()
		pos := yPosition
		if (yPosition > len(line)) {
			pos =yPosition % len(line)

		}

		if (string(line[pos]) == tree) {
			treeCount++
		}
		yPosition += 3
		lineCount += 1

	}

	fmt.Println(treeCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


}
