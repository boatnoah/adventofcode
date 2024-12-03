package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/boatnoah/aoc-go/common"
)

func partOne() {
	contents, err := common.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	re := regexp.MustCompile(`mul\(([0-9]+),\s*([0-9]+)\)`)

	var product int

	for _, char := range contents {
		matches := re.FindAllStringSubmatch(char, -1)

		for _, match := range matches {
			x, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Printf("%v", err)
				return
			}

			y, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Printf("%v", err)
				return
			}

			product += mul(x, y)
		}

	}

	fmt.Println(product)
}

func partTwo() {
	contents, err := common.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	re := regexp.MustCompile(`mul\(([0-9]+),\s*([0-9]+)\)|don\'t\(\)|do\(\)`)

	var product int

	multiplyFlag := true

	for _, char := range contents {
		matches := re.FindAllStringSubmatch(char, -1)

		for _, match := range matches {
			if strings.Contains(match[0], "don't()") {
				multiplyFlag = false
			} else if strings.Contains(match[0], "do()") {
				multiplyFlag = true
			} else {
				x, err := strconv.Atoi(match[1])
				if err != nil {
					fmt.Printf("%v", err)
					return
				}

				y, err := strconv.Atoi(match[2])
				if err != nil {
					fmt.Printf("%v", err)
					return
				}

				if multiplyFlag {
					product += mul(x, y)
				}

			}
		}

	}

	fmt.Println(product)
}

func mul(x int, y int) int {
	return x * y
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
