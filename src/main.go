package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	highest, lowest := scores[0], scores[0]
	result := []int32{0, 0}

	for _, value := range scores[1:] {
		if value > highest {
			highest = value
			result[0]++
		} else if value < lowest {
			lowest = value
			result[1]++
		}
	}

	return result
}

func main() {
	score := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	fmt.Println(breakingRecords(score))
}
