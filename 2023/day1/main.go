package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/boatnoah/aoc-go/common"
)

func main() {
	contents, err := common.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading files: %v", err)
		return
	}

	var sum int
	for i := range len(contents) {
		line := contents[i]

		var digits []rune

		for _, ch := range line {

			digit := ch

			if unicode.IsDigit(digit) {
				digits = append(digits, digit)
			}

		}

		joinedDigits := string(digits[0]) + string(digits[len(digits)-1])

		value, err := strconv.Atoi(joinedDigits)
		if err != nil {
			fmt.Println("Failed to parse string", err)
		}

		sum += value

	}

	fmt.Println(sum)
}
