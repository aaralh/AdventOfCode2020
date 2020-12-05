package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)

func assertEq(requiredFields []string, passportFields []string) bool {
	sort.Strings(requiredFields)
	sort.Strings(passportFields)
	fmt.Println("Required fields: ", requiredFields)
	fmt.Println("Current fields: ", passportFields)
	fmt.Println("")
	return reflect.DeepEqual(requiredFields, passportFields)
}

func main() {
	fmt.Println("Day 4")

	f, err := os.Open("./day4/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	requiredFields := []string{"byr", "iyr", "eyr" , "hgt", "hcl", "ecl", "pid"}
	var currentPassportFields []string

	validPassportsCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if (line == "") {
			if (assertEq(requiredFields, currentPassportFields)) {
				validPassportsCount++
			}
			currentPassportFields = nil
			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			fieldKey := strings.Split(field, ":")[0]
			if (fieldKey != "cid") {
				fmt.Println("Field: ", field)
				fmt.Println("Field key:", fieldKey)
				currentPassportFields = append(currentPassportFields, fieldKey)
			}
		}
	}

	if (assertEq(requiredFields, currentPassportFields)) {
		validPassportsCount++
	}

	fmt.Println("Valid passports: ", validPassportsCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


}