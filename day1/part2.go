package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Day 1")

	f, err := os.Open("./day1/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var numbers []int
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	fmt.Println(numbers)
	search:
	for _, number1 := range numbers {
		for _, number2 := range numbers {
			for _, number3 := range numbers {
				if number1+number2+number3 == 2020 {
					fmt.Println(number1 * number2 * number3)
					fmt.Println(number1, number2, number3)
					break search
				}
			}
		}
	}
}