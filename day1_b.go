package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	nums_key := make([]int, 1)
	nums_freq := make([]int, 1)
	for scanner.Scan() {
		a, b := getInputs(scanner.Text())
		nums_key = append(nums_key, a)
		nums_freq = append(nums_freq, b)
	}

	slices.Sort(nums_key)
	slices.Sort(nums_freq)

	var total int = 0
	for _, value := range nums_key {
		// fmt.Println("value:", value)
		for _, n := range nums_freq {
			// fmt.Println("\t compared to:", n)
			switch {
			case value < n:
				continue
			case value > n:
				break
			case value == n:
				total += value
				fmt.Println("\t adding, new total:", total)
			}
		}
	}
	fmt.Println("total:", total)
}

func getInputs(input string) (a int, b int) {
	fmt.Sscanf(input, "%d   %d", &a, &b)
	return a, b
}
