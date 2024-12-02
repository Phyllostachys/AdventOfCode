package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	firstList := make([]int32, 1)
	secondList := make([]int32, 1)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var firstItem, secondItem int32
		line := scanner.Text()
		fmt.Println(line)
		fmt.Sscanf(line, "%d   %d", &firstItem, &secondItem)
		fmt.Println(firstItem, secondItem)
		firstList = append(firstList, firstItem)
		secondList = append(secondList, secondItem)
	}
	slices.Sort(firstList)
	slices.Sort(secondList)
	var total int32 = 0
	for idx := 0; idx < len(firstList); idx++ {
		fmt.Println("first:", firstList[idx])
		fmt.Println("second:", secondList[idx])
		diff := firstList[idx] - secondList[idx]
		if diff < 0 {
			diff *= -1
		}
		fmt.Println("diff:", diff)
		total += diff
	}
	fmt.Println("total:", total)
}
