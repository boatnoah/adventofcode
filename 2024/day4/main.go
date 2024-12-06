package main

import (
	"fmt"

	"github.com/boatnoah/aoc-go/common"
)

type Pair struct {
	Row, Col int
}

type Set struct {
	elements map[Pair]bool
}

var directions = []Pair{
	{-1, 0},  // Up
	{1, 0},   // Down
	{0, -1},  // Left
	{0, 1},   // Right
	{-1, -1}, // Top-left
	{-1, 1},  // Top-right
	{1, -1},  // Bottom-left
	{1, 1},   // Bottom-right
}

var diagonalDirections = []Pair{
	{-1, -1}, // Top-left
	{-1, 1},  // Top-right
	{1, -1},  // Bottom-left
	{1, 1},   // Bottom-right
}

// NewSet creates a new set
func NewSet() *Set {
	return &Set{
		elements: make(map[Pair]bool),
	}
}

// Add adds an element to the set
func (s *Set) Add(element Pair) {
	s.elements[element] = true
}

// Remove removes an element from the set
func (s *Set) Remove(element Pair) {
	delete(s.elements, element)
}

// Contains checks if the set contains an element
func (s *Set) Contains(element Pair) bool {
	return s.elements[element]
}

// Size returns the size of the set
func (s *Set) Size() int {
	return len(s.elements)
}

// List returns all elements in the set
func (s *Set) List() []interface{} {
	keys := make([]interface{}, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func partOne() {
	contents, err := common.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Count not parse file")
		return
	}

	var grid [][]rune

	for _, line := range contents {
		var row []rune
		for _, chr := range line {
			row = append(row, chr)
		}
		grid = append(grid, row)
	}

	var appearance int

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			for _, direction := range directions {

				visited := NewSet()
				if dfs("XMAS", string(grid[row][col]), Pair{Row: row, Col: col}, grid, visited, direction) {
					appearance++
				}
			}
		}
	}

	fmt.Println(appearance)
}

func partTwo() {
	contents, err := common.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Count not parse file")
		return
	}

	var grid [][]rune

	for _, line := range contents {
		var row []rune
		for _, chr := range line {
			row = append(row, chr)
		}
		grid = append(grid, row)
	}

	var appearance int

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if isValidMas(Pair{Row: row, Col: col}, grid) {
				appearance++
			}
		}
	}

	fmt.Println(appearance)
}

func dfs(target string, word string, coordinates Pair, grid [][]rune, visited *Set, direction Pair) bool {
	// Base cases
	if len(word) > len(target) || coordinates.Row < 0 || coordinates.Row >= len(grid) || coordinates.Col < 0 || coordinates.Col >= len(grid[0]) || len(word) > len(target) {
		return false
	}

	if visited.Contains(coordinates) {
		return false
	}

	if word == target { // Found the target word
		return true
	}

	// Mark the current cell as visited
	visited.Add(coordinates)

	// Continue down that rabbit hole
	newRow := coordinates.Row + direction.Row
	newCol := coordinates.Col + direction.Col

	if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
		return false
	}

	newCoordinates := Pair{newRow, newCol}
	if dfs(target, word+string(grid[newRow][newCol]), newCoordinates, grid, visited, direction) {
		return true
	}

	// Backtrack: Unmark the current cell
	visited.Remove(coordinates)

	return false
}

func isValidMas(coordinates Pair, grid [][]rune) bool {
	if grid[coordinates.Row][coordinates.Col] != 'A' {
		return false // We only want to check 'A' positions
	}
	// Define the surrounding coordinates for the X-MAS pattern
	topLeft := Pair{Row: coordinates.Row - 1, Col: coordinates.Col - 1}
	topRight := Pair{Row: coordinates.Row - 1, Col: coordinates.Col + 1}
	bottomLeft := Pair{Row: coordinates.Row + 1, Col: coordinates.Col - 1}
	bottomRight := Pair{Row: coordinates.Row + 1, Col: coordinates.Col + 1}

	// Check if any of the coordinates are out of bounds
	if topLeft.Row < 0 || topLeft.Col < 0 || topRight.Row < 0 || topRight.Col >= len(grid[0]) ||
		bottomLeft.Row >= len(grid) || bottomLeft.Col < 0 || bottomRight.Row >= len(grid) || bottomRight.Col >= len(grid[0]) {
		return false // Out of bounds, cannot form the X shape
	}

	// Construct the "MAS" sequences for both parts of the X shape
	masOne := string(grid[topLeft.Row][topLeft.Col]) + string(grid[coordinates.Row][coordinates.Col]) + string(grid[bottomRight.Row][bottomRight.Col])
	masTwo := string(grid[topRight.Row][topRight.Col]) + string(grid[coordinates.Row][coordinates.Col]) + string(grid[bottomLeft.Row][bottomLeft.Col])

	m1 := freqArray(masOne)
	m2 := freqArray(masTwo)
	m3 := freqArray("MAS")

	return mapsAreEqual(m1, m3) && mapsAreEqual(m2, m3)
}

func mapsAreEqual(map1, map2 map[rune]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value := range map1 {
		if map2[key] != value {
			return false
		}
	}

	return true
}

func freqArray(s string) map[rune]int {
	freq := make(map[rune]int)

	for _, ch := range s {
		freq[ch]++
	}

	return freq
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
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
