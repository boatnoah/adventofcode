package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/boatnoah/aoc-go/common"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

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

	leftHeap := &IntHeap{}
	rightHeap := &IntHeap{}

	heap.Init(leftHeap)
	heap.Init(rightHeap)

	for _, ele := range contents {
		values := strings.Split(ele, " ")
		left, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Printf("Error casting string to int %v", err)
		}

		right, err := strconv.Atoi(values[len(values)-1])
		if err != nil {
			fmt.Printf("Error casting string to int %v", err)
		}

		heap.Push(leftHeap, left)
		heap.Push(rightHeap, right)
	}

	var result int

	for leftHeap.Len() > 0 && rightHeap.Len() > 0 {

		smallestLeft := heap.Pop(leftHeap).(int)
		smallestRight := heap.Pop(rightHeap).(int)

		result += diff(smallestLeft, smallestRight)
	}

	fmt.Println(result)
}

func partTwo() {
	contents, err := common.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error parsing file %v", err)
		return
	}

	leftMap := make(map[int]int)
	rightMap := make(map[int]int)

	for _, ele := range contents {
		values := strings.Split(ele, " ")
		leftValue, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Printf("Error casting string to int %v", err)
		}

		leftMap[leftValue]++

		rightValue, err := strconv.Atoi(values[len(values)-1])
		if err != nil {
			fmt.Printf("Error casting string to int %v", err)
		}

		rightMap[rightValue]++

	}

	var result int

	for k := range leftMap {
		if _, ok := rightMap[k]; ok {
			result += k * rightMap[k]
		}
	}

	fmt.Println(result)
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
