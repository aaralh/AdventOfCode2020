package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {
	fmt.Println("Day 3")

	var treeCounts []int

	slopes := [5][2]int{{1, 1}, {3, 1}, {5, 1},{7, 1}, {1, 2}}
	//slopes := [1][2]int{{1, 2}}

	for _, slope := range slopes {
		xMovement := slope[0]
		yMovement := slope[1]

		treeCount := 0
		lineCount := 0

		f, err := os.Open("./day3/data.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()
		scanner := bufio.NewScanner(f)
		tree := "#"
		xPosition := 0

		for scanner.Scan() {

			if (yMovement == 2 && lineCount%2 == 1) {
				lineCount += 1
				continue
			}

			line := scanner.Text()
			pos := xPosition
			if (xPosition >= len(line)) {
				pos = xPosition % len(line)
			}

			if (string(line[pos]) == tree) {
				treeCount++
			}
			xPosition += xMovement
			lineCount += 1

		}
		fmt.Print(treeCount, "*")
		treeCounts = append(treeCounts, treeCount)


		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(treeCounts)

}
