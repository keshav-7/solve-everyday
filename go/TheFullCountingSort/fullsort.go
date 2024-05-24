package main

import (
	"fmt"
	"strconv"
)

func countSort(arr [][]string) {
	// Find the maximum value from the first elements
	maxVal := 0
	for _, v := range arr {
		num, _ := strconv.Atoi(v[0])
		if num > maxVal {
			maxVal = num
		}
	}

	// Create a slice of slices to hold the counts
	count := make([][]string, maxVal+1)

	// Replace strings in the first half of the array with '-'
	for i := range arr {
		num, _ := strconv.Atoi(arr[i][0])
		if i < len(arr)/2 {
			count[num] = append(count[num], "-")
		} else {
			count[num] = append(count[num], arr[i][1])
		}
	}

	// Flatten the count slice and print the sorted strings
	for _, sublist := range count {
		for _, str := range sublist {
			fmt.Print(str + " ")
		}
	}
	fmt.Println()
}

func main() {
	// Sample Input
	arr := [][]string{
		{"0", "ab"},
		{"6", "cd"},
		{"0", "ef"},
		{"6", "gh"},
		{"4", "ij"},
		{"0", "ab"},
		{"6", "cd"},
		{"0", "ef"},
		{"6", "gh"},
		{"0", "ij"},
		{"4", "that"},
		{"3", "be"},
		{"0", "to"},
		{"1", "be"},
		{"5", "question"},
		{"1", "or"},
		{"2", "not"},
		{"4", "is"},
		{"2", "to"},
		{"4", "the"},
	}

	countSort(arr)
}
