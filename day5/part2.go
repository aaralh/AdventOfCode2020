package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("./day5/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var ids []string

	for scanner.Scan() {
		line := scanner.Text()
		ids = append(ids, line)
	}

	sort.Strings(ids)


	var values []int


	for _, code := range ids {
		rowMin := 0
		rowMax := 127
		colMin := 0
		colMax := 7

		for index, char := range code {
			if (index < 7) {
				if (string(char) == "B") {
					if (rowMax - rowMin == 1) {
						rowMin = rowMax
					} else {
						rowMin += ((rowMax + 1) - rowMin)/2
					}
				} else if (string(char) == "F") {
					if (rowMax - rowMin == 1) {
						rowMax = rowMin
					} else {
						rowMax -= (rowMax + 1 - rowMin) / 2
					}
				}
			} else {
				if (string(char) == "R") {
					colMin += ((colMax + 1) - colMin)/2
				} else if (string(char) == "L") {
					colMax -= ((colMax +  1) - colMin)/2
				}
			}
		}
		val := (rowMin * 8) + colMin
		values = append(values, val)
	}
	sort.Ints(values)
	prev := 0

	for index, val := range values {
		if (index == 0) {
			prev = val
			continue
		}
		if (val - prev == 2) {
			fmt.Println("Values: ", prev, val)
		}
		prev = val
	}
}
