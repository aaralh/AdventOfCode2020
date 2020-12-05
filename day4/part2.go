package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
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
	return reflect.DeepEqual(requiredFields, passportFields)
}

func isValidLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'f') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

func validateData(passportData map[string]string, passportFields []string) bool {
	dataIsValid := true
	for _, field := range passportFields {
		data := passportData[field]
		switch field {
			case "byr":
				if (len(data) != 4) {
					dataIsValid = false
					break
				}
				year, err := strconv.Atoi(data)

				if err != nil {
					dataIsValid = false
				}

				if (1920 > year || year > 2002 ) {
					dataIsValid = false
				}
				fmt.Println(field, dataIsValid)
				break
			case "iyr":
				if (len(data) != 4) {
					dataIsValid = false
					break
				}
				year, err := strconv.Atoi(data)

				if err != nil {
					dataIsValid = false
				}

				if (2010 > year || year > 2020 ) {
					dataIsValid = false
				}
				fmt.Println(field, dataIsValid)
				break

			case "eyr":

				if (len(data) != 4) {
					dataIsValid = false
					break
				}

				year, err := strconv.Atoi(data)

				if err != nil {
					dataIsValid = false
				}

				if (2020 > year || year > 2030 ) {
					dataIsValid = false
				}
				fmt.Println(field, dataIsValid)
				break

			case "hgt":
				if strings.HasSuffix(data, "cm") {
					height, err := strconv.Atoi(strings.Trim(data, "cm"))
					if err != nil {
						dataIsValid = false
					}

					if (150 > height || height > 193 ) {
						dataIsValid = false
					}

				} else if strings.HasSuffix(data, "in") {
					height, err := strconv.Atoi(strings.Trim(data, "in"))
					if err != nil {
						dataIsValid = false
					}

					if (59 > height || height > 76 ) {
						dataIsValid = false
					}
				} else {
					dataIsValid = false
				}
				fmt.Println(field, dataIsValid)
				break

			case "hcl":
				if (!strings.HasPrefix(data, "#")) {
					dataIsValid = false
				}

				if (len(data) != 7) {
					dataIsValid = false
				}

				if (!isValidLetter(strings.Trim(data, "#"))) {
					dataIsValid = false
				}
				fmt.Println(field, dataIsValid)
				break

			case "ecl":
				switch data {
					case
						"amb", "blu", "brn", "gry", "grn", "hzl", "oth":
							break
					default:
						dataIsValid = false
					}
				fmt.Println(field, dataIsValid)
				break

			case "pid":
				if (len(data) != 9) {
					dataIsValid = false
				} else {
					isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
					if strings.IndexFunc(data, isNotDigit) == 0 {
						dataIsValid = false
					}
				}
				fmt.Println(field, dataIsValid)
				break

			default:
				// freebsd, openbsd,
				// plan9, windows...
				fmt.Printf("Here we are")
		}
	}
	fmt.Println(dataIsValid)
	return dataIsValid
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
	currentPassportData := make(map[string]string)

	validPassportsCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if (line == "") {
			if (assertEq(requiredFields, currentPassportFields)) {
				fmt.Println(currentPassportData)

				if (validateData(currentPassportData, currentPassportFields)) {
					validPassportsCount++
				}
				fmt.Println("")
			}
			currentPassportFields = nil
			currentPassportData = make(map[string]string)
			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			fieldKey := strings.Split(field, ":")[0]
			fieldValue := strings.Split(field, ":")[1]
			if (fieldKey != "cid") {
				currentPassportData[fieldKey] = fieldValue
				currentPassportFields = append(currentPassportFields, fieldKey)
			}
		}
	}

	if (assertEq(requiredFields, currentPassportFields)) {
		if (validateData(currentPassportData, currentPassportFields)) {
			validPassportsCount++
		}
	}

	fmt.Println("Valid passports: ", validPassportsCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


}