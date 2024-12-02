package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/boatnoah/aoc-go/common"
)

func diff(a int, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

func partOne() {
	contents, err := common.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error parsing file %v", err)
		return
	}

	var safeCount int

	for _, char := range contents {
		res := strings.Split(char, " ")

		var digits []int

		for _, char := range res {
			digit, err := strconv.Atoi(char)
			if err != nil {
				fmt.Printf("Error parsing digit, %v", err)
				return
			}
			digits = append(digits, digit)
		}

		if isSafe(digits) {
			safeCount++
		}

	}

	fmt.Println(safeCount)
}

func partTwo() {
	contents, err := common.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error parsing file %v", err)
		return
	}

	var safeCount int

	for _, char := range contents {
		res := strings.Split(char, " ")

		var digits []int

		for _, char := range res {
			digit, err := strconv.Atoi(char)
			if err != nil {
				fmt.Printf("Error parsing digit, %v", err)
				return
			}
			digits = append(digits, digit)
		}

		safeFlag := true
		prev := digits[0]
		i := 1

		ascendingFlag := (prev < digits[1]) // true is ascending false is descending

		for i < len(digits) {

			curr := digits[i]
			absDiff := diff(curr, prev)

			if (absDiff > 3) || (absDiff < 1) || (ascendingFlag != (prev < curr)) {

				var local bool

				var j int

				for j < len(digits) {
					newSlice := make([]int, len(digits))
					copy(newSlice, digits)

					slices.Delete(newSlice, j, j+1)
					newSlice = newSlice[:len(newSlice)-1]

					if isSafe(newSlice) {
						local = true
						break
					}
					j++

				}

				if !local {
					safeFlag = false
					break
				}

			}

			prev = curr
			i++

		}

		if safeFlag {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func isSafe(digits []int) bool {
	prev := digits[0]
	i := 1

	ascendingFlag := (prev < digits[1]) // true is ascending false is descending

	for i < len(digits) {

		curr := digits[i]
		absDiff := diff(curr, prev)

		if (absDiff > 3) || (absDiff < 1) || (ascendingFlag != (prev < curr)) {
			return false
		}

		prev = curr
		i++

	}

	return true
}

func main() {
	var input string
	fmt.Print("Show Part 1 or 2: ")
	fmt.Scanln(&input)

	if input == "1" {
		partOne()
	} else {
		partTwo()
	}
}
